package render

import (
	"log"
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

// TestRenderCommentsWithQuotes addresses https://github.com/pacedotdev/oto/issues/17.
func TestRenderCommentsWithQuotes(t *testing.T) {
	is := is.New(t)
	def := parser.Definition{
		PackageName: "services",
		Services: []parser.Service{
			{
				Comment: `This comment contains "quotes"`,
				Name:    "MyService",
			},
		},
	}
	template := `
		<%= for (service) in def.Services { %>
			<%= format_comment_text(service.Comment) %>type <%= service.Name %> struct
		<% } %>
	`
	s, err := Render(template, def, nil)
	is.NoErr(err)
	log.Println(s)
	for _, should := range []string{
		`// This comment contains "quotes"`,
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

func TestFormatCommentText(t *testing.T) {
	is := is.New(t)

	actual := strings.TrimSpace(string(formatCommentText("card's")))
	is.Equal(actual, "// card's")

	actual = strings.TrimSpace(string(formatCommentText(`What happens if I use "quotes"?`)))
	is.Equal(actual, `// What happens if I use "quotes"?`)

	actual = strings.TrimSpace(string(formatCommentText("What about\nnew lines?")))
	is.Equal(actual, `// What about new lines?`)

}

func TestFormatComment(t *testing.T) {
	is := is.New(t)

	longComment := `This is a long comment that will end up spanning
multiple lines so we get to test the indent string option
in formatComment.`
	actual := strings.TrimSpace(string(formatComment(longComment, "\t\t")))
	is.Equal(actual, `This is a long comment that will end up spanning multiple lines so we get to
		test the indent string option in formatComment.`)

}
