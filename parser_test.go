package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestParse(t *testing.T) {
	is := is.New(t)
	patterns := []string{"./testdata/services/pleasantries"}
	def, err := newParser(patterns...).parse()
	is.NoErr(err)

	is.Equal(def.PackageName, "pleasantries")
	is.Equal(len(def.Services), 2)
	is.Equal(def.Services[0].Name, "GreeterService")
	is.Equal(len(def.Services[0].Methods), 2)
	is.Equal(def.Services[0].Methods[0].Name, "GetGreetings")
	is.Equal(def.Services[0].Methods[0].InputObject.TypeName, "GetGreetingsRequest")
	is.Equal(def.Services[0].Methods[0].InputObject.Multiple, false)
	is.Equal(def.Services[0].Methods[0].InputObject.Package, "")
	is.Equal(def.Services[0].Methods[0].OutputObject.TypeName, "GetGreetingsResponse")
	is.Equal(def.Services[0].Methods[0].OutputObject.Multiple, false)
	is.Equal(def.Services[0].Methods[0].OutputObject.Package, "")

	is.Equal(def.Services[0].Methods[1].Name, "Greet")
	is.Equal(def.Services[0].Methods[1].InputObject.TypeName, "GreetRequest")
	is.Equal(def.Services[0].Methods[1].InputObject.Multiple, false)
	is.Equal(def.Services[0].Methods[1].InputObject.Package, "")
	is.Equal(def.Services[0].Methods[1].OutputObject.TypeName, "GreetResponse")
	is.Equal(def.Services[0].Methods[1].OutputObject.Multiple, false)
	is.Equal(def.Services[0].Methods[1].OutputObject.Package, "")

	greetInputObject, err := def.Object(def.Services[0].Methods[0].InputObject.TypeName)
	is.NoErr(err)
	is.Equal(greetInputObject.Name, "GetGreetingsRequest")
	is.Equal(len(greetInputObject.Fields), 1)
	is.Equal(greetInputObject.Fields[0].Name, "Page")
	is.Equal(greetInputObject.Fields[0].OmitEmpty, false)
	is.Equal(greetInputObject.Fields[0].Type.TypeName, "services.Page")
	is.Equal(greetInputObject.Fields[0].Type.TypeID, "github.com/pacelabs/oto/testdata/services.Page")
	is.Equal(greetInputObject.Fields[0].Type.IsObject, true)
	is.Equal(greetInputObject.Fields[0].Type.Multiple, false)
	is.Equal(greetInputObject.Fields[0].Type.Package, "github.com/pacelabs/oto/testdata/services")

	greetOutputObject, err := def.Object(def.Services[0].Methods[0].OutputObject.TypeName)
	is.NoErr(err)
	is.Equal(greetOutputObject.Name, "GetGreetingsResponse")
	is.Equal(len(greetOutputObject.Fields), 2)
	is.Equal(greetOutputObject.Fields[0].Name, "Greetings")
	is.Equal(greetOutputObject.Fields[0].Type.TypeID, "github.com/pacelabs/oto/testdata/services/pleasantries.Greeting")
	is.Equal(greetOutputObject.Fields[0].OmitEmpty, false)
	is.Equal(greetOutputObject.Fields[0].Type.TypeName, "Greeting")
	is.Equal(greetOutputObject.Fields[0].Type.Multiple, true)
	is.Equal(greetOutputObject.Fields[0].Type.Package, "")
	is.Equal(greetOutputObject.Fields[1].Name, "Error")
	is.Equal(greetOutputObject.Fields[1].OmitEmpty, true)
	is.Equal(greetOutputObject.Fields[1].Type.TypeName, "string")
	is.Equal(greetOutputObject.Fields[1].Type.Multiple, false)
	is.Equal(greetOutputObject.Fields[1].Type.Package, "")

	is.Equal(def.Services[1].Name, "Welcomer")
	is.Equal(len(def.Services[1].Methods), 1)

	is.Equal(def.Services[1].Methods[0].InputObject.TypeName, "WelcomeRequest")
	is.Equal(def.Services[1].Methods[0].InputObject.Multiple, false)
	is.Equal(def.Services[1].Methods[0].InputObject.Package, "")
	is.Equal(def.Services[1].Methods[0].OutputObject.TypeName, "WelcomeResponse")
	is.Equal(def.Services[1].Methods[0].OutputObject.Multiple, false)
	is.Equal(def.Services[1].Methods[0].OutputObject.Package, "")

	welcomeInputObject, err := def.Object(def.Services[1].Methods[0].InputObject.TypeName)
	is.NoErr(err)
	is.Equal(welcomeInputObject.Name, "WelcomeRequest")
	is.Equal(len(welcomeInputObject.Fields), 2)
	is.Equal(welcomeInputObject.Fields[0].Name, "To")
	is.Equal(welcomeInputObject.Fields[0].OmitEmpty, false)
	is.Equal(welcomeInputObject.Fields[0].Type.TypeName, "string")
	is.Equal(welcomeInputObject.Fields[0].Type.Multiple, false)
	is.Equal(welcomeInputObject.Fields[0].Type.Package, "")
	is.Equal(welcomeInputObject.Fields[1].Name, "Name")
	is.Equal(welcomeInputObject.Fields[1].OmitEmpty, false)
	is.Equal(welcomeInputObject.Fields[1].Type.TypeName, "string")
	is.Equal(welcomeInputObject.Fields[1].Type.Multiple, false)
	is.Equal(welcomeInputObject.Fields[1].Type.Package, "")

	welcomeOutputObject, err := def.Object(def.Services[1].Methods[0].OutputObject.TypeName)
	is.NoErr(err)
	is.Equal(welcomeOutputObject.Name, "WelcomeResponse")
	is.Equal(len(welcomeOutputObject.Fields), 2)
	is.Equal(welcomeOutputObject.Fields[0].Name, "Message")
	is.Equal(welcomeOutputObject.Fields[0].Type.IsObject, false)
	is.Equal(welcomeOutputObject.Fields[0].OmitEmpty, false)
	is.Equal(welcomeOutputObject.Fields[0].Type.TypeName, "string")
	is.Equal(welcomeOutputObject.Fields[0].Type.Multiple, false)
	is.Equal(welcomeOutputObject.Fields[0].Type.Package, "")
	is.Equal(welcomeOutputObject.Fields[1].Name, "Error")
	is.Equal(welcomeOutputObject.Fields[1].OmitEmpty, true)
	is.Equal(welcomeOutputObject.Fields[1].Type.TypeName, "string")
	is.Equal(welcomeOutputObject.Fields[1].Type.Multiple, false)
	is.Equal(welcomeOutputObject.Fields[1].Type.Package, "")

	is.Equal(len(def.Objects), 8)
	for i := range def.Objects {
		switch def.Objects[i].Name {
		case "Greeting":
			is.Equal(len(def.Objects[i].Fields), 1)
			is.Equal(def.Objects[i].Imported, false)
		case "Page":
			is.Equal(def.Objects[i].TypeID, "github.com/pacelabs/oto/testdata/services.Page")
			is.Equal(len(def.Objects[i].Fields), 3)
			is.Equal(def.Objects[i].Imported, true)
		}
	}

	// b, err := json.MarshalIndent(def, "", "  ")
	// is.NoErr(err)
	// log.Println(string(b))
}

func TestFieldJSType(t *testing.T) {
	is := is.New(t)
	for in, expected := range map[fieldType]string{
		fieldType{TypeName: "string"}:                     "string",
		fieldType{TypeName: "int"}:                        "number",
		fieldType{TypeName: "uint"}:                       "number",
		fieldType{TypeName: "uint32"}:                     "number",
		fieldType{TypeName: "int32"}:                      "number",
		fieldType{TypeName: "int64"}:                      "number",
		fieldType{TypeName: "float64"}:                    "number",
		fieldType{TypeName: "bool"}:                       "boolean",
		fieldType{TypeName: "interface{}"}:                "any",
		fieldType{TypeName: "map[string]interface{}"}:     "object",
		fieldType{TypeName: "SomeObject", IsObject: true}: "object",
	} {
		actual, err := in.JSType()
		is.NoErr(err)
		if actual != expected {
			t.Errorf("%s expected: %q but got %q", in.TypeName, expected, actual)
		}
	}
}
