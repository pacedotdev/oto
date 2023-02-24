package render

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	"github.com/pacedotdev/oto/parser"
)

func ObjectGolang(def parser.Definition, object *parser.Object, tabs int) template.HTML {
	s := &strings.Builder{}
	fmt.Fprintf(s, "%s{", object.ExternalObjectName)
	for _, field := range object.Fields {
		fmt.Fprintf(s, "\n")
		fmt.Fprint(s, strings.Repeat("\t", tabs+1))
		if field.Type.IsObject {
			// object
			fieldObject, err := def.Object(field.Type.ObjectName)
			if err != nil {
				return template.HTML(field.Type.ObjectName + ": " + err.Error())
			}
			fmt.Fprintf(s, "%s: %v,", field.Name, ObjectGolang(def, fieldObject, tabs+1))
			continue
		}
		// normal field
		fmt.Fprintf(s, "%s: %v,", field.Name, jsonStr(field.Metadata["example"]))
	}
	fmt.Fprintf(s, "\n")
	fmt.Fprint(s, strings.Repeat("\t", tabs))
	fmt.Fprintf(s, "}")
	return template.HTML(s.String())
}

func jsonStr(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
