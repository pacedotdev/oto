package render

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/pacedotdev/oto/parser"
)

func TestRender(t *testing.T) {
	is := is.New(t)
	def := parser.Definition{
		PackageName: "services",
	}
	params := map[string]interface{}{
		"Description": "Package services contains services.",
	}
	template := `// <%= params["Description"] %>
package <%= def.PackageName %>`
	s, err := Render(template, def, params)
	is.NoErr(err)
	for _, should := range []string{
		"// Package services contains services.",
		"package services",
	} {
		if !strings.Contains(s, should) {
			t.Errorf("missing: %s", should)
			is.Fail()
		}
	}
}

func TestCamelizeDown(t *testing.T) {
	for in, expected := range map[string]string{
		"CamelsAreGreat": "camelsAreGreat",
		"ID":             "id",
		"HTML":           "html",
		"PreviewHTML":    "previewHTML",
	} {
		actual := camelizeDown(in)
		if actual != expected {
			t.Errorf("%s expected: %q but got %q", in, expected, actual)
		}
	}
}

func TestFormatTags(t *testing.T) {
	is := is.New(t)

	trimBackticks := func(s string) string {
		is.True(strings.HasPrefix(s, "`"))
		is.True(strings.HasSuffix(s, "`"))
		return strings.Trim(s, "`")
	}

	tagStr, err := formatTags(`json:"field,omitempty"`)
	is.NoErr(err)
	is.Equal(trimBackticks(string(tagStr)), `json:"field,omitempty"`)

	tagStr, err = formatTags(`json:"field,omitempty" monkey:"true"`)
	is.NoErr(err)
	is.Equal(trimBackticks(string(tagStr)), `json:"field,omitempty" monkey:"true"`)

	tagStr, err = formatTags(`json:"field,omitempty"`, `monkey:"true"`)
	is.NoErr(err)
	is.Equal(trimBackticks(string(tagStr)), `json:"field,omitempty" monkey:"true"`)

}
