package xsonschema

import (
	"encoding/json"

	"github.com/globalsign/mgo/bson"
)

type (
	Object struct{}

	StringMaybeList []string

	SchemaMaybeList []Schema

	SchemaListOrBoolean struct {
		Schemas SchemaMaybeList
		Boolean bool
	}

	Schema struct {
		// All Types

		// Accepts same string aliases used for the $type operator
		// string alias or array of string aliases
		BSONType StringMaybeList `json:"bsonType,omitempty" bson:"bsonType,omitempty"`

		// Enumerates the possible JSON types of the field. Available types are “object”, “array”, “number”, “boolean”, “string”, and “null”.
		// string or array of unique strings
		Type StringMaybeList `json:"type,omitempty" bson:"type,omitempty"`

		// Enumerates all possible values of the field
		Enum []interface{} `json:"enum,omitempty" bson:"enum,omitempty"`

		// Field must match all specified schemas
		AllOf []Schema `json:"allOf,omitempty" bson:"allOf,omitempty"`

		// Field must match at least one of the specified schemas
		AnyOf []Schema `json:"anyOf,omitempty" bson:"anyOf,omitempty"`

		// Field must match exactly one of the specified schemas
		OneOf []Schema `json:"oneOf,omitempty" bson:"oneOf,omitempty"`

		// Field must not match the schema
		// a JSON Schema object
		Not *Schema `json:"allOf,omitempty" bson:"allOf,omitempty"`

		// Numbers

		// Field must be a multiple of this value
		MultipleOf float64 `json:"multipleOf,omitempty" bson:"multipleOf,omitempty"`

		// Indicates the maximum value of the field
		Maximum float64 `json:"maximum,omitempty" bson:"maximum,omitempty"`

		// If true and field is a number, maximum is an exclusive maximum. Otherwise, it is an inclusive maximum.
		ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty" bson:"exclusiveMaximum,omitempty"`

		// Indicates the minimum value of the field
		Minimum float64 `json:"minimum,omitempty" bson:"minimum,omitempty"`

		// If true, minimum is an exclusive minimum. Otherwise, it is an inclusive minimum.
		ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty" bson:"exclusiveMinimum,omitempty"`

		// Array

		// Indicates the maximum length of array
		MaxItems int `json:"maxItems,omitempty" bson:"maxItems,omitempty"`

		// Indicates the minimum length of array
		MinItems int `json:"minItems,omitempty" bson:"minItems,omitempty"`

		// If true, each item in the array must be unique. Otherwise, no uniqueness constraint is enforced.
		UniqueItems bool `json:"uniqueItems,omitempty" bson:"uniqueItems,omitempty"`

		// If an object, must be a valid JSON Schema
		AdditionalItems SchemaListOrBoolean `json:"additionalItems,omitempty" bson:"additionalItems,omitempty"`

		// Must be either a valid JSON Schema, or an array of valid JSON Schemas
		Items SchemaMaybeList `json:"items,omitempty" bson:"items,omitempty"`

		// strings

		// Indicates the maximum length of the field

		MaxLength int `json:"maxLength,omitempty" bson:"maxLength,omitempty"`
		// Indicates the minimum length of the field

		MinLength int `json:"minLength,omitempty" bson:"minLength,omitempty"`

		// Field must match the regular expression
		Pattern string `json:"pattern,omitempty" bson:"pattern,omitempty"`

		// objects

		// Indicates the field’s maximum number of properties
		MaxProperties int `json:"maxProperties,omitempty" bson:"maxProperties,omitempty"`

		// Indicates the field’s minimum number of properties
		MinProperties int `json:"minProperties,omitempty" bson:"minProperties,omitempty"`

		// Object’s property set must contain all the specified elements in the array
		Required []string `json:"required,omitempty" bson:"required,omitempty"`

		// If true, additional fields are allowed. If false, they are not.
		// If a valid JSON Schema object is specified, additional fields must validate against the schema. Defaults to true.
		AdditionalProperties SchemaListOrBoolean `json:"additionalProperties,omitempty" bson:"additionalProperties,omitempty"`

		// A valid JSON Schema where each value is also a valid JSON Schema object
		Properties map[string]Schema `json:"properties,omitempty" bson:"properties,omitempty"`

		// In addition to properties requirements, each property name of this object must be a valid regular expression
		// objects	object
		PatternProperties map[string]Schema `json:"patternProperties,omitempty" bson:"patternProperties,omitempty"`

		// Describes field or schema dependencies
		// objects	object
		Dependencies struct{} `json:"dependencies,omitempty" bson:"dependencies,omitempty"`
	}
)

func Parse(schema Schema, data []byte) (Object, error) {

	return Object{}, nil
}

func (list StringMaybeList) MarshalJSON() ([]byte, error) {
	if len(list) < 2 {
		return json.Marshal(list[0])
	}
	return json.Marshal(list)
}

func (list *StringMaybeList) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		(*list) = append((*list), str)
	} else {
		var strs []string
		if err := json.Unmarshal(data, &strs); err != nil {
			return err
		}
		(*list) = StringMaybeList(strs)
	}
	return nil
}
func (list StringMaybeList) MarshalBSON() ([]byte, error) {
	if len(list) < 2 {
		return bson.Marshal(list[0])
	}
	return bson.Marshal(list)
}
func (list *StringMaybeList) UnmarshalBSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var str string
		if err := bson.Unmarshal(data, &str); err != nil {
			return err
		}
		(*list) = append((*list), str)
	} else {
		var strs []string
		if err := bson.Unmarshal(data, &strs); err != nil {
			return err
		}
		(*list) = StringMaybeList(strs)
	}
	return nil
}

func (list SchemaMaybeList) MarshalJSON() ([]byte, error) {
	if len(list) < 2 {
		return json.Marshal(list[0])
	}
	return json.Marshal([]Schema(list))
}

func (list *SchemaMaybeList) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var schema Schema
		if err := json.Unmarshal(data, &schema); err != nil {
			return err
		}
		(*list) = append((*list), schema)
	} else {
		var schemaList []Schema
		if err := json.Unmarshal(data, &schemaList); err != nil {
			return err
		}
		(*list) = SchemaMaybeList(schemaList)
	}
	return nil
}
func (list SchemaMaybeList) MarshalBSON() ([]byte, error) {
	if len(list) < 2 {
		return bson.Marshal(list[0])
	}
	return bson.Marshal([]Schema(list))
}
func (list *SchemaMaybeList) UnmarshalBSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var schema Schema
		if err := bson.Unmarshal(data, &schema); err != nil {
			return err
		}
		(*list) = append((*list), schema)
	} else {
		var schemaList []Schema
		if err := bson.Unmarshal(data, &schemaList); err != nil {
			return err
		}
		(*list) = SchemaMaybeList(schemaList)
	}
	return nil
}
func (item SchemaListOrBoolean) MarshalJSON() ([]byte, error) {
	if len(item.Schemas) == 0 {
		return json.Marshal(item.Boolean)
	}
	return json.Marshal(item.Schemas)
}
func (list *SchemaListOrBoolean) UnmarshalJSON(data []byte) error {
	if data[0] == '[' || data[0] == '{' {
		return json.Unmarshal(data, &list.Schemas)
	} else {
		return json.Unmarshal(data, &list.Boolean)
	}
}
func (item SchemaListOrBoolean) MarshalBSON() ([]byte, error) {
	if len(item.Schemas) == 0 {
		return bson.Marshal(item.Boolean)
	}
	return bson.Marshal(item.Schemas)
}
func (list *SchemaListOrBoolean) UnmarshalBSON(data []byte) error {
	if data[0] == '[' || data[0] == '{' {
		return bson.Unmarshal(data, &list.Schemas)
	} else {
		return bson.Unmarshal(data, &list.Boolean)
	}
}

// func () MarshalJSON() ([]byte, error){ return []byte{}, nil }
// func () UnmarshalJSON([]byte) error{ return nil }
// func () MarshalBSON() ([]byte, error){ return []byte{}, nil }
// func () UnmarshalBSON([]byte) error{ return nil }
