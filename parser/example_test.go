package parser

import (
	"testing"

	"github.com/matryer/is"
)

func TestObjectExample(t *testing.T) {
	is := is.New(t)

	obj1 := Object{
		Name: "obj1",
		Fields: []Field{
			{
				Name:           "Name",
				NameLowerCamel: "name",
				Example:        "Mat",
			},
			{
				Name:           "Project",
				NameLowerCamel: "project",
				Example:        "Respond",
			},
			{
				Name:           "SinceYear",
				NameLowerCamel: "sinceYear",
				Example:        2021,
			},
			{
				Name:           "Favourites",
				NameLowerCamel: "favourites",
				Type: FieldType{
					TypeName:        "obj2",
					IsObject:        true,
					CleanObjectName: "obj2",
				},
			},
			{
				Name:           "SomeObject",
				NameLowerCamel: "someObject",
				Type: FieldType{
					TypeName:        "obj3",
					IsObject:        true,
					CleanObjectName: "obj3",
					Multiple:        false,
				},
			},
			{
				Name:           "Tags",
				NameLowerCamel: "tags",
				Type: FieldType{
					Multiple: true,
					TypeName: "string",
				},
				Example: []interface{}{"security", "customer-affected", "review-needed"},
			},
		},
	}
	obj2 := Object{
		Name: "obj2",
		Fields: []Field{
			{
				Type:           FieldType{TypeName: "string", Multiple: true, CleanObjectName: "string"},
				NameLowerCamel: "languages",
				Example:        []interface{}{"Go"},
			},
		},
	}
	obj3 := Object{
		Name: "obj3",
		Fields: []Field{
			{
				Name:           "Tags",
				NameLowerCamel: "tags",
				Example:        []interface{}{"one", "two", "three"},
				Type: FieldType{
					TypeName:        "obj2",
					CleanObjectName: "obj2",
					Multiple:        true,
					IsObject:        true,
				},
			},
		},
	}
	def := &Definition{
		Objects: []Object{obj1, obj2, obj3},
	}
	example, err := def.Example(obj1)
	is.NoErr(err)
	is.True(example != nil)

	// check it out:
	// b, err := json.MarshalIndent(example, "", "  ")
	// is.NoErr(err)
	// fmt.Println(string(b))

	is.Equal(example["name"], "Mat")
	is.Equal(example["project"], "Respond")
	is.Equal(example["sinceYear"], 2021)
	is.True(example["favourites"] != nil)
	favourites, ok := example["favourites"].(map[string]interface{})
	is.True(ok) // Favourites map[string]interface{}
	languages, ok := favourites["languages"].([]interface{})
	is.True(ok) // Languages []interface{}
	is.Equal(len(languages), 1)

	exampleJSON, err := def.Example(obj1)
	is.NoErr(err)

	is.Equal(len(exampleJSON), 6)
	is.Equal(len(exampleJSON["tags"].([]interface{})), 3)
	is.Equal(exampleJSON["tags"].([]interface{})[0], "security")

}
