package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func Test(t *testing.T) {
	is := is.New(t)
	var buf bytes.Buffer
	args := []string{
		"oto",
		"-template=./testdata/template.plush",
		"-pkg=stuff",
		"./testdata/services/pleasantries",
	}
	err := run(&buf, args)
	is.NoErr(err)
	s := buf.String()
	for _, should := range []string{
		"GreeterService.GetGreetings",
		"GreeterService.Greet",
		"Welcomer.Welcome",
	} {
		if !strings.Contains(s, should) {
			t.Errorf("missing: %s", should)
			is.Fail()
		}
	}
}

func TestParseParams(t *testing.T) {
	is := is.New(t)

	params, err := parseParams("key1:value1,key2: value2 , key3:value3")
	is.NoErr(err)
	is.Equal(params["key1"], "value1")
	is.Equal(params["key2"], "value2")
	is.Equal(params["key3"], "value3")

}
