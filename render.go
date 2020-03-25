package main

import (
	"github.com/gobuffalo/plush"
	"github.com/markbates/inflect"
)

var defaultRuleset = inflect.NewDefaultRuleset()

// render renders the template using the definition.
func render(template string, def definition, params map[string]interface{}) (string, error) {
	ctx := plush.NewContext()
	ctx.Set("camelize_down", camelizeDown)
	ctx.Set("underscore", underscore)
	ctx.Set("def", def)
	ctx.Set("params", params)
	ctx.Set("rust_type", rustType)
	ctx.Set("has", has)
	s, err := plush.Render(string(template), ctx)
	if err != nil {
		return "", err
	}
	return s, nil
}

// camelizeDown converts a name or other string into a camel case
// version with the first letter lowercase. "ModelID" becomes "modelID".
func camelizeDown(s string) string {
	if s == "ID" {
		return "id"
		// note: not sure why I need this, there's a lot that deals with
		// accronyms in the dependency packages but they don't seem to behave
		// as expected in this case.
	}
	return defaultRuleset.CamelizeDownFirst(s)
}

// underscore converts a name or other string into a snake case
// version. "ModelID" becomes "model_id".
func underscore(s string) string {
	if s == "ID" {
		return "id"
	}
	return defaultRuleset.Underscore(s)
}

// rustType converst the given type name to its rust equivalent
func rustType(s string) string {
	switch s {
	case "string":
		return "String"
	case "int":
		return "i64"
	case "int8":
		return "i8"
	case "int16":
		return "i16"
	case "int32":
		return "i32"
	case "int64":
		return "i64"
	case "uint":
		return "u64"
	case "uint8":
		return "u8"
	case "uint16":
		return "u16"
	case "uint32":
		return "u32"
	case "uint64":
		return "u64"
	case "float32":
		return "f32"
	case "float64":
		return "f64"
	case "map[string]interface{}":
		return "Object"
	default:
		return s
	}
}

// camelizeDown converts a name or other string into a camel case
// version with the first letter lowercase. "ModelID" becomes "modelID".
func has(input object, fieldName string) bool {
	for _, field := range input.Fields {
		if field.Name == fieldName {
			return true
		}
	}
	return false
}
