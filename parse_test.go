package xsoncheck_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/crhntr/xsoncheck"
	"github.com/globalsign/mgo/bson"
)

func Test(t *testing.T) {

	tests := map[string]bool{
		"testdata/schema_00.json": false,
		"testdata/schema_01.json": true,
		"testdata/schema_02.json": true,
	}

	for testFileName, shouldErr := range tests {
		t.Run(testFileName, func(t *testing.T) {
			f, err := os.Open(testFileName)
			if err != nil {
				t.Fatal(err)
			}
			var schema xsoncheck.Schema
			if err := json.NewDecoder(f).Decode(&schema); !shouldErr && err != nil {
				t.Error(err)
			} else if shouldErr && err == nil {
				t.Error("err should have occured")
			}

			if !t.Failed() {
				buf, err := bson.Marshal(schema)
				if err != nil {
					t.Error(err)
				}

				var schema xsoncheck.Schema
				err = bson.Unmarshal(buf, &schema)
				if err != nil {
					t.Error(err)
				}
			}

		})
	}
}

func EqualSlices(strs1 []string, strs2 ...string) bool {
	if len(strs1) != len(strs2) {
		return false
	}
	for i, str := range strs1 {
		if str == strs2[i] {
			return false
		}
	}

	return true
}
