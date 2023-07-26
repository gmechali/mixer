// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package writer

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
	"github.com/datacommonsorg/mixer/internal/server/resource"
	"github.com/datacommonsorg/mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var triplesInitData = []*triple{
	{subjectID: "dc/g/New", predicate: "typeOf", objectID: "StatVarGroup"},
	{subjectID: "dc/g/New", predicate: "name", objectValue: "New Variables"},
	{subjectID: "dc/g/New", predicate: "specializationOf", objectID: "dc/g/Root"},
}

type observation struct {
	entity     string
	variable   string
	date       string
	value      string
	provenance string
}

type triple struct {
	subjectID   string
	predicate   string
	objectID    string
	objectValue string
}

// Write writes raw CSV files to SQLite CSV files.
func Write(resourceMetadata *resource.Metadata) error {
	fileDir := resourceMetadata.SQLitePath
	csvFiles, err := listCSVFiles(fileDir)
	if err != nil {
		return err
	}
	if len(csvFiles) == 0 {
		return status.Errorf(codes.FailedPrecondition, "No CSV files found in %s", fileDir)
	}
	observationList := []*observation{}
	tripleList := triplesInitData
	variableSet := map[string]struct{}{}
	for _, csvFile := range csvFiles {
		provID := fmt.Sprintf("dc/custom/%s", strings.TrimRight(csvFile, ".csv"))
		observations, variables, err := processCSVFile(
			resourceMetadata, fileDir, csvFile, provID)
		if err != nil {
			return err
		}
		observationList = append(observationList, observations...)
		for _, v := range variables {
			variableSet[v] = struct{}{}
		}
		tripleList = append(tripleList,
			&triple{
				subjectID: provID,
				predicate: "dcid",
				objectID:  provID,
			},
			&triple{
				subjectID: provID,
				predicate: "typeOf",
				objectID:  "Provenance",
			},
			&triple{
				subjectID:   provID,
				predicate:   "url",
				objectValue: filepath.Join(fileDir, csvFile),
			},
		)
	}
	for variable := range variableSet {
		tripleList = append(tripleList,
			&triple{
				subjectID: variable,
				predicate: "typeOf",
				objectID:  "StatisticalVariable",
			},
			&triple{
				subjectID: variable,
				predicate: "memberOf",
				objectID:  "dc/g/New",
			},
			&triple{
				subjectID:   variable,
				predicate:   "description",
				objectValue: variable,
			},
		)
	}

	return writeOutput(observationList, tripleList, fileDir)
}

func listCSVFiles(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, file := range files {
		if fName := file.Name(); strings.HasSuffix(fName, ".csv") {
			res = append(res, file.Name())
		}
	}

	return res, nil
}

func processCSVFile(
	medatata *resource.Metadata,
	fileDir string,
	csvFile string,
	provID string,
) (
	[]*observation,
	[]string, // A list of variables.
	error,
) {
	// Read the CSV file.
	f, err := os.Open(filepath.Join(fileDir, csvFile))
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, nil, err
	}
	numRecords := len(records)
	if numRecords < 2 {
		return nil, nil, status.Errorf(codes.FailedPrecondition,
			"Empty CSV file %s", csvFile)
	}

	// Load header.
	header := records[0]
	if len(header) < 3 {
		return nil, nil, status.Errorf(codes.FailedPrecondition,
			"Less than 3 columns in CSV file %s", csvFile)
	}
	numColumns := len(header)

	// Resolve places.
	places := []string{}
	for i := 1; i < numRecords; i++ {
		places = append(places, records[i][0])
	}
	resolvedPlaceMap, err := resolvePlaces(medatata, places, header[0])
	if err != nil {
		return nil, nil, err
	}

	// Generate observations.
	observations := []*observation{}
	for i := 1; i < numRecords; i++ {
		record := records[i]

		resolvedPlace, ok := resolvedPlaceMap[record[0]]
		if !ok {
			// If a place cannot be resolved, simply ignore it.
			continue
		}

		for j := 2; j < numColumns; j++ {
			observations = append(observations, &observation{
				entity:     resolvedPlace,
				variable:   header[j],
				date:       record[1],
				value:      record[j],
				provenance: provID,
			})
		}
	}

	return observations, header[2:], nil
}

func resolvePlaces(
	metadata *resource.Metadata,
	places []string,
	placeHeader string,
) (map[string]string, error) {
	placeToDCID := map[string]string{}
	if placeHeader == "lat#lng" {
		// TODO(ws): lat#lng recon.
	} else if placeHeader == "name" {
		// TODO(ws): name recon.
	} else {
		resp := &pbv2.ResolveResponse{}
		httpClient := &http.Client{}
		if err := util.FetchRemote(metadata, httpClient, "/v2/resolve",
			&pbv2.ResolveRequest{
				Nodes:    places,
				Property: fmt.Sprintf("<-%s->dcid", placeHeader),
			}, resp); err != nil {
			return nil, err
		}
		for _, entity := range resp.GetEntities() {
			if _, ok := placeToDCID[entity.GetNode()]; ok {
				continue
			}
			// TODO(ws): Handle the case with multiple ResolvedIds.
			placeToDCID[entity.GetNode()] = entity.GetCandidates()[0].GetDcid()
		}
	}

	return placeToDCID, nil
}

func prepareDatabase(fileDir string) error {
	dbPath := filepath.Join(fileDir, "datacommons.db")
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		_, err := os.Create(dbPath)
		if err != nil {
			return err
		}
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`
			DROP TABLE IF EXISTS observations;
			DROP TABLE IF EXISTS triples;
		`,
	)
	if err != nil {
		return err
	}

	tripleStatement := `
	CREATE TABLE triples (
		subject_id TEXT,
		predicate TEXT,
		object_id TEXT,
		object_value TEXT
	);
	`
	_, err = db.Exec(tripleStatement)
	if err != nil {
		return err
	}

	observationStatement := `
	CREATE TABLE observations (
		entity TEXT,
		variable TEXT,
		date TEXT,
		value TEXT,
		provenance TEXT
	);
	`
	_, err = db.Exec(observationStatement)
	if err != nil {
		return err
	}
	return nil
}

func writeOutput(
	observations []*observation,
	triples []*triple,
	fileDir string,
) error {
	err := prepareDatabase(fileDir)
	if err != nil {
		return err
	}
	db, err := sql.Open("sqlite3", filepath.Join(fileDir, "datacommons.db"))
	if err != nil {
		return err
	}
	defer db.Close()

	// Observations.
	for _, o := range observations {
		sqlStmt := `INSERT INTO observations(entity,variable,date,value,provenance) VALUES (?, ?, ?, ?, ?)`
		_, err = db.Exec(sqlStmt, o.entity, o.variable, o.date, o.value, o.provenance)
		if err != nil {
			return err
		}
	}

	// Triples.
	for _, t := range triples {
		sqlStmt := `INSERT INTO triples(subject_id,predicate,object_id,object_value) VALUES (?, ?, ?, ?)`
		_, err = db.Exec(sqlStmt, t.subjectID, t.predicate, t.objectID, t.objectValue)
		if err != nil {
			return err
		}
	}
	return nil
}