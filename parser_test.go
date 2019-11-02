package oto

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
	b, err := json.MarshalIndent(def, "", "  ")
	is.NoErr(err)
	log.Printf("%+v\n", string(b))

}
