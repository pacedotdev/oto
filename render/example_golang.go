package render

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/pacedotdev/oto/parser"
)

func ObjectGolang(def parser.Definition, object parser.Object, tabs int) string {
	s := &strings.Builder{}
	fmt.Fprintf(s, strings.Repeat("\t", tabs))
	for _, field := range object.Fields {
		fmt.Fprintf(s, "\n")
		fmt.Fprintf(s, strings.Repeat("\t", tabs+1))
		if field.Type.IsObject {
			// object
			fieldObject, err := def.Object(field.Type.ObjectName)
			if err != nil {
				return field.Type.ObjectName + ": " + err.Error()
			}
			fmt.Fprintf(s, `%s: %s{
%v
}`, field.Name, field.Type.TypeName, ObjectGolang(def, *fieldObject, tabs+2))
			continue
		}
		// normal field
		log.Printf("%+v", field)
		fmt.Fprintf(s, "%s: %v", field.Name, jsonStr(field.Metadata["example"]))
	}
	fmt.Fprintf(s, "\n")
	fmt.Fprintf(s, strings.Repeat("\t", tabs))
	return s.String()
}

func jsonStr(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
