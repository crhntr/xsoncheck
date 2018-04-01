package xsonschema

type (
	Object struct{}

	SchemaMaybeList []Schema

	SchemaListOrBoolean struct {
		Schema  SchemaMaybeList
		Boolean bool
	}

	Schema struct {
		// All Types

		// Accepts same string aliases used for the $type operator
		// string alias or array of string aliases
		BSONType []string `json:"bsonType,omitempty" bson:"bsonType,omitempty"`

		// Enumerates the possible JSON types of the field. Available types are “object”, “array”, “number”, “boolean”, “string”, and “null”.
		// string or array of unique strings
		Type []string `json:"type,omitempty" bson:"type,omitempty"`

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
		MultipleOf float64

		// Indicates the maximum value of the field
		Maximum float64

		// If true and field is a number, maximum is an exclusive maximum. Otherwise, it is an inclusive maximum.
		ExclusiveMaximum bool

		// Indicates the minimum value of the field
		Minimum float64

		// If true, minimum is an exclusive minimum. Otherwise, it is an inclusive minimum.
		ExclusiveMinimum bool

		// Array

		// Indicates the maximum length of array
		MaxItems int `json:"maxItems,omitempty" bson:"maxItems,omitempty"`

		// Indicates the minimum length of array
		MinItems int `json:"minItems,omitempty" bson:"minItems,omitempty"`

		// If true, each item in the array must be unique. Otherwise, no uniqueness constraint is enforced.
		//boolean
		UniqueItems bool `json:"uniqueItems,omitempty" bson:"uniqueItems,omitempty"`

		// If an object, must be a valid JSON Schema
		// might be bool
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

/*




title	// N/A	string	A descriptive title string with no effect.
description	// N/A	string	A string that describes the schema and has no effect.
*/
