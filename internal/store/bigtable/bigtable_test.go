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

package bigtable

import (
	"context"
	"testing"

	cbt "cloud.google.com/go/bigtable"
	"github.com/google/go-cmp/cmp"
)

func TestReadOneTable(t *testing.T) {
	ctx := context.Background()
	data := map[string]string{
		"key1": "data1",
		"key2": "data2",
	}
	btTable, err := SetupBigtable(ctx, data)
	if err != nil {
		t.Errorf("setupBigtable got error: %v", err)
	}
	rowList := cbt.RowList{"key1", "key2"}
	btData, err := Read(
		ctx,
		NewGroup([]*Table{{name: "test", table: btTable}}, ""),
		rowList,
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			return string(jsonRaw), nil
		},
		nil,
	)
	if err != nil {
		t.Errorf("btReadRowsParallel got error: %v", err)
	}
	for dcid, result := range btData[0] {
		if diff := cmp.Diff(data[dcid], result.(string)); diff != "" {
			t.Errorf("read rows got diff from table data %+v", diff)
		}
	}
}

func TestReadTwoTables(t *testing.T) {
	ctx := context.Background()

	data1 := map[string]string{
		"key1": "foo1",
		"key2": "foo2",
	}
	data2 := map[string]string{
		"key1": "bar1",
		"key2": "bar2",
		"key3": "bar3",
	}

	btTable1, err := SetupBigtable(ctx, data1)
	if err != nil {
		t.Errorf("setupBigtable1 got error: %v", err)
	}

	btTable2, err := SetupBigtable(ctx, data2)
	if err != nil {
		t.Errorf("setupBigtable2 got error: %v", err)
	}

	rowList := cbt.RowList{"key1", "key2"}
	dataList, err := Read(
		ctx,
		NewGroup(
			[]*Table{
				{name: "t1_t", table: btTable1},
				{name: "t2_t", table: btTable2},
			},
			"t2",
		),
		rowList,
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			return string(jsonRaw), nil
		},
		nil,
	)
	if err != nil {
		t.Errorf("btReadRowsParallel got error: %v", err)
	}
	for dcid, result := range dataList[0] {
		if diff := cmp.Diff(data1[dcid], result.(string)); diff != "" {
			t.Errorf("read rows got diff from table data %+v", diff)
		}
	}
}
