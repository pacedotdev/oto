package main

import (
	"github.com/gobuffalo/plush"
)

// render renders the template using the definition.
func render(template string, def definition) (string, error) {
	ctx := plush.NewContext()
	ctx.Set("def", def)
	s, err := plush.Render(string(template), ctx)
	if err != nil {
		return "", err
	}
	return s, nil
}
