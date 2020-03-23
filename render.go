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
