package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestRender(t *testing.T) {
	is := is.New(t)
	def := definition{
		PackageName: "services",
	}
	params := map[string]interface{}{
		"Description": "Package services contains services.",
	}
	template := `// <%= params["Description"] %>
package <%= def.PackageName %>`
	s, err := render(template, def, params)
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
	} {
		actual := camelizeDown(in)
		if actual != expected {
			t.Errorf("%s expected: %q but got %q", in, expected, actual)
		}
	}
}

func TestStructTag2(t *testing.T) {
	for expected, in := range map[string][]string{
		"":                            {"", ""},
		"`db:\"name\" json:\"name\"`": {`db:"name" json:"name"`, ""},
		"`db:\"name\" json:\"foo\" fruit:\"apple\"`": {`db:"name" json:"name"`, `json:"foo" fruit:"apple"`},
		"`json:\"foo\"`": {"", `json:"foo"`},
	} {
		actual := string(structTag2(in[0], in[1]))

		if actual != expected {
			t.Errorf("%s expected: %q but got %q", in, expected, actual)
		}
	}
}
