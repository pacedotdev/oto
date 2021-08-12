package parser

import (
	"encoding/json"
	"html/template"
)

// Example generates an object that is a realistic example
// of this object.
// Examples are read from the docs.
// This is experimental.
func (d *Definition) Example(o Object) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	for _, field := range o.Fields {
		if field.Type.IsObject {
			subobj, err := d.Object(field.Type.TypeName)
			if err != nil {
				return nil, err
			}
			example, err := d.Example(*subobj)
			if err != nil {
				return nil, err
			}
			obj[field.NameLowerCamel] = example
			if field.Type.Multiple {
				// turn it into an array
				obj[field.NameLowerCamel] = []interface{}{obj[field.NameLowerCamel]}
			}
			continue
		}
		obj[field.NameLowerCamel] = field.Example
		if field.Type.Multiple {
			// turn it into an array
			obj[field.NameLowerCamel] = []interface{}{obj[field.NameLowerCamel], obj[field.NameLowerCamel], obj[field.NameLowerCamel]}
		}
	}
	return obj, nil
}

// ExampleJSON is like Example, but returns a JSON string.
func (d *Definition) ExampleJSON(o Object) (template.HTML, error) {
	example, err := d.Example(o)
	if err != nil {
		return "", err
	}
	exampleBytes, err := json.MarshalIndent(example, "", "\t")
	if err != nil {
		return "", err
	}
	return template.HTML(exampleBytes), nil
}
