package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestRender(t *testing.T) {
	is := is.New(t)
	def := definition{
		PackageName: "services",
	}
	template := `package <%= def.PackageName %>`
	s, err := render(template, def)
	is.NoErr(err)
	is.Equal(s, `package services`)
}
