package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/matryer/is"
)

func TestParse(t *testing.T) {
	is := is.New(t)
	patterns := []string{"./testdata/services/pleasantries"}
	def, err := newParser(patterns...).parse()
	is.NoErr(err)
	is.Equal(len(def.Services), 2)
	// TODO: more assertions once the design has settled

	b, err := json.MarshalIndent(def, "", "  ")
	is.NoErr(err)
	log.Println(string(b))
}
