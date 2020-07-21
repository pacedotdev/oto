package main

import (
	"bytes"
	"encoding/json"
	"go/doc"
	"html/template"

	"github.com/fatih/structtag"
	"github.com/gobuffalo/plush"
	"github.com/markbates/inflect"
	"github.com/pkg/errors"
)

var defaultRuleset = inflect.NewDefaultRuleset()

// render renders the template using the Definition.
func render(template string, def Definition, params map[string]interface{}) (string, error) {
	ctx := plush.NewContext()
	ctx.Set("camelize_down", camelizeDown)
	ctx.Set("def", def)
	ctx.Set("params", params)
	ctx.Set("json", toJSONHelper)
	ctx.Set("format_comment_text", formatCommentText)
	ctx.Set("format_comment_html", formatCommentHTML)
	ctx.Set("format_tags", formatTags)
	s, err := plush.Render(string(template), ctx)
	if err != nil {
		return "", err
	}
	return s, nil
}

func toJSONHelper(v interface{}) (template.HTML, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.HTML(b), nil
}

func formatCommentText(s string) string {
	var buf bytes.Buffer
	doc.ToText(&buf, s, "// ", "", 80)
	return buf.String()
}

func formatCommentHTML(s string) template.HTML {
	var buf bytes.Buffer
	doc.ToHTML(&buf, s, nil)
	return template.HTML(buf.String())
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

// formatTags formats a list of struct tag strings into one.
// Will return an error if any of the tag strings are invalid.
func formatTags(tags ...string) (template.HTML, error) {
	alltags := &structtag.Tags{}
	for _, tag := range tags {
		theseTags, err := structtag.Parse(tag)
		if err != nil {
			return "", errors.Wrapf(err, "parse tags: `%s`", tag)
		}
		for _, t := range theseTags.Tags() {
			alltags.Set(t)
		}
	}
	tagsStr := alltags.String()
	if tagsStr == "" {
		return "", nil
	}
	tagsStr = "`" + tagsStr + "`"
	return template.HTML(tagsStr), nil
}
