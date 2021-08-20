// Copyright 2020 Google LLC
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

package server

import (
	"context"
	"log"
	"sort"
	"strings"
	"time"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetStatValue implements API for Mixer.GetStatValue.
// Endpoint: /stat (/stat/value)
func (s *Server) GetStatValue(ctx context.Context, in *pb.GetStatValueRequest) (
	*pb.GetStatValueResponse, error) {
	place := in.GetPlace()
	statVar := in.GetStatVar()

	if place == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: place")
	}
	if statVar == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: stat_var")
	}
	date := in.GetDate()
	filterProp := &ObsProp{
		Mmethod: in.GetMeasurementMethod(),
		Operiod: in.GetObservationPeriod(),
		Unit:    in.GetUnit(),
		Sfactor: in.GetScalingFactor(),
	}

	rowList, keyTokens := buildStatsKey([]string{place}, []string{statVar})
	var obsTimeSeries *ObsTimeSeries
	btData, err := readStats(ctx, s.store, rowList, keyTokens)
	if err != nil {
		return nil, err
	}
	obsTimeSeries = btData[place][statVar]
	if obsTimeSeries == nil {
		return nil, status.Errorf(
			codes.NotFound, "No data for %s, %s", place, statVar)
	}
	obsTimeSeries.SourceSeries = filterSeries(obsTimeSeries.SourceSeries, filterProp)
	result, err := getValueFromBestSource(obsTimeSeries, date)
	if err != nil {
		return nil, err
	}
	return &pb.GetStatValueResponse{Value: result}, nil
}

func getStatSet(
	ctx context.Context, s *Server, places []string, statVars []string, date string) (
	*pb.GetStatSetResponse, error) {
	// Initialize result with stat vars and place dcids.
	ts := time.Now()
	result := &pb.GetStatSetResponse{
		Data: make(map[string]*pb.PlacePointStat),
	}
	for _, statVar := range statVars {
		result.Data[statVar] = &pb.PlacePointStat{
			Stat:     make(map[string]*pb.PointStat),
			Metadata: make(map[string]*pb.StatMetadata),
		}
		for _, place := range places {
			result.Data[statVar].Stat[place] = nil
		}
	}

	rowList, keyTokens := buildStatsKey(places, statVars)
	cacheData, err := readStatsPb(ctx, s.store, rowList, keyTokens)
	if err != nil {
		return nil, err
	}
	for place, placeData := range cacheData {
		for statVar, data := range placeData {
			if data != nil {
				stat, meta := getValueFromBestSourcePb(data, date)
				result.Data[statVar].Stat[place] = stat
				if meta != nil {
					result.Data[statVar].Metadata[meta.ImportName] = meta
				}
			}
		}
	}
	log.Printf("getStatSet() completed for %d places, %d stat vars, in %s seconds",
		len(places), len(statVars), time.Since(ts))
	return result, nil
}

// GetStatSet implements API for Mixer.GetStatSet.
// Endpoint: /stat/set
func (s *Server) GetStatSet(ctx context.Context, in *pb.GetStatSetRequest) (
	*pb.GetStatSetResponse, error) {
	places := in.GetPlaces()
	statVars := in.GetStatVars()
	date := in.GetDate()

	if len(places) == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: places")
	}
	if len(statVars) == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: stat_vars")
	}
	return getStatSet(ctx, s, places, statVars, date)
}

// GetStatSetWithinPlace implements API for Mixer.GetStatSetWithinPlace.
// Endpoint: /stat/set/within-place
func (s *Server) GetStatSetWithinPlace(
	ctx context.Context, in *pb.GetStatSetWithinPlaceRequest) (
	*pb.GetStatSetResponse, error) {
	parentPlace := in.GetParentPlace()
	statVars := in.GetStatVars()
	childType := in.GetChildType()
	date := in.GetDate()

	log.Printf(
		"GetStatSetWithinPlace: parentPlace: %s, statVars: %v, childType: %s, date: %s",
		parentPlace,
		statVars,
		childType,
		date,
	)
	if parentPlace == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: parent_place")
	}
	if len(statVars) == 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: stat_vars")
	}
	if childType == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"Missing required argument: child_type")
	}
	if date == "" {
		// Try to read from cache directly
		rowList, keyTokens := buildStatSetWithinPlaceKey(parentPlace, childType, "LATEST", statVars)
		cacheData, err := readStatCollection(ctx, s.store, rowList, keyTokens)
		if err != nil {
			return nil, err
		}
		result := &pb.GetStatSetResponse{
			Data: make(map[string]*pb.PlacePointStat),
		}
		for _, statVar := range statVars {
			result.Data[statVar] = &pb.PlacePointStat{
				Stat:     make(map[string]*pb.PointStat),
				Metadata: make(map[string]*pb.StatMetadata),
			}
		}
		gotResult := false
		for statVar, data := range cacheData {
			if data != nil {
				gotResult = true
				cohorts := data.SourceCohorts
				// Sort cohort first, so the preferred source is populated first. Then only
				// update when there is a later data.
				sort.Sort(SeriesByRank(cohorts))
				for _, cohort := range cohorts {
					for place, val := range cohort.Val {
						pointStat, ok := result.Data[statVar].Stat[place]
						if !ok || pointStat.Date < cohort.PlaceToLatestDate[place] {
							metaData := &pb.StatMetadata{
								MeasurementMethod: cohort.MeasurementMethod,
								ObservationPeriod: cohort.ObservationPeriod,
								ProvenanceUrl:     cohort.ProvenanceUrl,
								ScalingFactor:     cohort.ScalingFactor,
								ImportName:        cohort.ImportName,
							}
							result.Data[statVar].Stat[place] = &pb.PointStat{
								Date:     cohort.PlaceToLatestDate[place],
								Value:    val,
								Metadata: metaData,
							}
							result.Data[statVar].Metadata[cohort.ImportName] = metaData
						}
					}
				}
			}
		}
		if gotResult {
			return result, nil
		}
	}

	// No data found from cache, fetch stat series for each place separately.
	// Get all the child places
	rowList := buildPlaceInKey([]string{parentPlace}, childType)
	// Place relations are from base geo imports. Only trust the base cache.
	baseDataMap, _, err := bigTableReadRowsParallel(
		ctx,
		s.store,
		rowList,
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			return strings.Split(string(jsonRaw), ","), nil
		},
		nil,
		false, /* readBranch */
	)
	if err != nil {
		return nil, err
	}
	if baseDataMap[parentPlace] == nil {
		return &pb.GetStatSetResponse{
			Data: make(map[string]*pb.PlacePointStat),
		}, nil
	}
	childPlaces := baseDataMap[parentPlace].([]string)
	return getStatSet(ctx, s, childPlaces, statVars, date)
}
