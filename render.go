package main

import (
	"html/template"

	"github.com/fatih/structtag"
	"github.com/gobuffalo/plush"
	"github.com/markbates/inflect"
)

var defaultRuleset = inflect.NewDefaultRuleset()

// render renders the template using the definition.
func render(template string, def definition, params map[string]interface{}) (string, error) {
	ctx := plush.NewContext()
	ctx.Set("camelize_down", camelizeDown)
	ctx.Set("def", def)
	ctx.Set("params", params)
	ctx.Set("struct_tag", structTag)
	ctx.Set("struct_tag2", structTag2)
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

func backticks(s string) string {
	return "`" + s + "`"
}

var emptyHTML = template.HTML("")

func structTag(tag string) template.HTML {
	return structTag2(tag, "")
}

func structTag2(tagstr string, additional string) template.HTML {
	if tagstr != "" && additional != "" {
		tags, err := structtag.Parse(tagstr)
		if err != nil {
			return emptyHTML
		}

		add, err := structtag.Parse(additional)
		if err != nil {
			return template.HTML(backticks(tags.String()))
		}

		for _, item := range add.Tags() {
			tags.Set(item)
		}
		return template.HTML(backticks(tags.String()))
	} else if tagstr != "" {
		return template.HTML(backticks(tagstr))
	} else if additional != "" {
		return template.HTML(backticks(additional))
	}

	return emptyHTML
}
