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
