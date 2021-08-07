package parser

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestObjectExample(t *testing.T) {
	is := is.New(t)

	obj1 := Object{
		Name: "obj1",
		Fields: []Field{
			{
				Name:    "Name",
				Example: "Mat",
			},
			{
				Name:    "Project",
				Example: "Respond",
			},
			{
				Name:    "SinceYear",
				Example: 2021,
			},
			{
				Name: "Favourites",
				Type: FieldType{
					TypeName: "obj2",
					IsObject: true,
				},
			},
		},
	}
	obj2 := Object{
		Name: "obj2",
		Fields: []Field{
			{
				Type:    FieldType{TypeName: "string", Multiple: true},
				Name:    "Languages",
				Example: "Go",
			},
		},
	}
	def := &Definition{
		Objects: []Object{obj1, obj2},
	}
	example, err := def.Example(obj1)
	is.NoErr(err)
	is.True(example != nil)

	b, err := json.MarshalIndent(example, "", "\t")
	is.NoErr(err)
	fmt.Println(string(b))

	is.Equal(example["Name"], "Mat")
	is.Equal(example["Project"], "Respond")
	is.Equal(example["SinceYear"], 2021)
	is.True(example["Favourites"] != nil)
	favourites, ok := example["Favourites"].(map[string]interface{})
	is.True(ok) // Favourites map[string]interface{}
	languages, ok := favourites["Languages"].([]interface{})
	is.True(ok) // Languages []interface{}
	is.Equal(len(languages), 3)

}
