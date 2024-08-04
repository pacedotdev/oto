package render

import (
	"bytes"
	"encoding/json"
	"go/doc"
	"html/template"
	"strings"

	"github.com/fatih/structtag"
	"github.com/gobuffalo/plush"
	"github.com/pacedotdev/oto/parser"
	"github.com/pkg/errors"
)

// Render renders the template using the Definition.
func Render(template string, def parser.Definition, params map[string]interface{}) (string, error) {
	ctx := plush.NewContext()
	ctx.Set("camelize_down", camelizeDown)
	ctx.Set("camelize_up", camelizeUp)
	ctx.Set("camelize_up_field", camelizeUpField)
	ctx.Set("def", def)
	ctx.Set("params", params)
	ctx.Set("json", toJSONHelper)
	ctx.Set("json_inline", toJSONInlineHelper)
	ctx.Set("format_comment_line", formatCommentLine)
	ctx.Set("format_comment_text", formatCommentText)
	ctx.Set("format_comment_html", formatCommentHTML)
	ctx.Set("format_tags", formatTags)
	ctx.Set("object_golang", ObjectGolang)
	ctx.Set("smart_prefix", smartPrefix)
	s, err := plush.Render(string(template), ctx)
	if err != nil {
		return "", err
	}
	return s, nil
}

func toJSONHelper(v interface{}, prefix, indent string) (template.HTML, error) {
	if indent == "" {
		indent = "\t"
	}
	b, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return "", err
	}
	return template.HTML(b), nil
}

func toJSONInlineHelper(v interface{}) (template.HTML, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.HTML(b), nil
}

func formatCommentLine(s string) template.HTML {
	var buf bytes.Buffer
	doc.ToText(&buf, s, "", "", 2000)
	s = strings.TrimSpace(buf.String())
	return template.HTML(s)
}

func formatCommentText(s string) template.HTML {
	var buf bytes.Buffer
	doc.ToText(&buf, s, "// ", "", 80)
	return template.HTML(buf.String())
}

func formatCommentHTML(s string) template.HTML {
	var buf bytes.Buffer
	doc.ToHTML(&buf, s, nil)
	return template.HTML(buf.String())
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

// smartPrefix prepends a string before s, allowing for the specific use
// case of pointers to objects. If the s begins with * (as in, *Object), the
// result will be *prefixObject to preserve its original meaning.
func smartPrefix(prefix, s string) string {
	if strings.HasPrefix(s, "*") {
		return "*" + prefix + s[1:]
	}
	return prefix + s
}
