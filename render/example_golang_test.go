package render

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/pacedotdev/oto/parser"
)

func TestExmapleGolang(t *testing.T) {
	is := is.New(t)
	patterns := []string{"../parser/testdata/services/pleasantries"}
	parser := parser.New(patterns...)
	parser.Verbose = testing.Verbose()
	parser.ExcludeInterfaces = []string{"Ignorer"}
	def, err := parser.Parse()
	is.NoErr(err)
	inputObject, err := def.Object("GreetRequest")
	is.NoErr(err) // get inputObject
	example := ObjectGolang(def, inputObject, 0)

	f, err := os.Create("./delete-me-example.go.notgo")
	is.NoErr(err) // write file
	defer f.Close()
	f.Write([]byte(example))

	log.Printf("\n---\n %s \n---\n", example)

	for _, should := range []string{
		"GetGreetingsRequest{",
		"Page: Page{",
	} {
		if !strings.Contains(string(example), should) {
			t.Errorf("missing: %s", should)
			is.Fail()
		}
	}
}
