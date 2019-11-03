package main

import (
	"bytes"
	"log"
	"testing"

	"github.com/matryer/is"
)

func Test(t *testing.T) {
	is := is.New(t)
	var buf bytes.Buffer
	args := []string{
		"oto",
		"-template=./testdata/templates/server.go.plush",
		"-pkg=stuff",
		"./testdata/services/pleasantries",
	}
	err := run(&buf, args)
	is.NoErr(err)
	log.Println(buf.String())
}
