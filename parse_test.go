package xsonschema_test

import (
	"encoding/json"
	"testing"

	"github.com/crhntr/xsonschema"
)

func Test(t *testing.T) {
	var schema xsonschema.Schema
	if err := json.Unmarshal([]byte(`{
    "bsonType": "object",
    "required": [ "name", "year", "major", "gpa" ],
    "properties": {
      "name": {
        "bsonType": "string",
        "description": "must be a string and is required"
      },
      "gender": {
        "bsonType": "string",
        "description": "must be a string and is not required"
      },
      "year": {
        "bsonType": "int",
        "minimum": 2017,
        "maximum": 3017,
        "exclusiveMaximum": false,
        "description": "must be an integer in [ 2017, 3017 ] and is required"
      },
      "major": {
        "enum": [ "Math", "English", "Computer Science", "History", null ],
        "description": "can only be one of the enum values and is required"
      },
      "gpa": {
        "bsonType": [ "double" ],
        "minimum": 0,
        "description": "must be a double and is required"
      }
    }
  }`), &schema); err != nil {
		t.Fatal(err)
	}

	if EqualSlices(schema.BSONType, "object") ||
		EqualSlices(schema.Required, "name", "year", "major", "gpa") {
		t.Fail()
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

/*

 */
