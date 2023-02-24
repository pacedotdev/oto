package render

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/pacedotdev/oto/parser"
)

func ObjectGolang(def parser.Definition, parentField *parser.Field, object *parser.Object, tabs int) template.HTML {
	s := &strings.Builder{}
	fmt.Fprintf(s, strings.Repeat("\t", tabs))
	if parentField != nil {
		fmt.Fprintf(s, "%s{\b", parentField.Type.ExternalObjectName)
	}
	for _, field := range object.Fields {
		fmt.Fprintf(s, "\n")
		fmt.Fprintf(s, strings.Repeat("\t", tabs+1))
		if field.Type.IsObject {
			// object
			fieldObject, err := def.Object(field.Type.ObjectName)
			if err != nil {
				return template.HTML(field.Type.ObjectName + ": " + err.Error())
			}
			fmt.Fprintf(s, "%s: %s{\n%v", field.Name, field.Type.ExternalObjectName, ObjectGolang(def, &field, fieldObject, tabs+1))
			fmt.Fprintf(s, "\n")
			fmt.Fprintf(s, strings.Repeat("\t", tabs+1))
			fmt.Fprintf(s, "},")
			continue
		}
		// normal field
		log.Printf("%+v", field)
		fmt.Fprintf(s, "%s: %v,", field.Name, jsonStr(field.Metadata["example"]))
	}
	fmt.Fprintf(s, strings.Repeat("\t", tabs+1))
	if parentField != nil {
		fmt.Fprintf(s, "}")
	}
	return template.HTML(s.String())
}

func jsonStr(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
