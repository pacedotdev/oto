package main

import (
	"fmt"
	"html/template"
	"strings"

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
	ctx.Set("rust_default", rustDefault)
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

// rustDefault returns the rust code for the default value
func rustDefault(s string) template.HTML {
	switch s {
	case "string":
		return "String::new()"
	case "int":
		fallthrough
	case "int8":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		fallthrough
	case "uint":
		fallthrough
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		return "std::default::Default::default()"
	case "float32":
		fallthrough
	case "float64":
		return "0.0"
	default:
		if strings.HasPrefix(s, "map[") {
			keyType := rustType(s[4 : 4+strings.Index(s[4:], "]")])
			valueType := rustType(s[5+len(keyType):])
			return template.HTML(fmt.Sprintf("std::collections::HashMap::<%s, %s>::new()", keyType, valueType))
		} else if strings.HasPrefix(s, "[]") {
			return template.HTML(fmt.Sprintf("std::Vec::new()"))
		}
		return template.HTML("Default::default()")
	}
}

// rustType converst the given type name to its rust equivalent
func rustType(s string) template.HTML {
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
	case "interface{}":
		return "Value"
	default:
		if strings.HasPrefix(s, "map[") {
			keyType := rustType(s[4 : 4+strings.Index(s[4:], "]")])
			valueType := rustType(s[5+len(keyType):])
			//return template.HTML("Map<String, Value>")
			return template.HTML(fmt.Sprintf("std::collections::HashMap<%s, %s>", keyType, valueType))
		}
		return template.HTML(s)
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
