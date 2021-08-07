package parser

import "encoding/json"

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
			obj[field.Name] = example
			if field.Type.Multiple {
				// turn it into an array
				obj[field.Name] = []interface{}{obj[field.Name]}
			}
			continue
		}
		obj[field.Name] = field.Example
		if field.Type.Multiple {
			// turn it into an array
			obj[field.Name] = []interface{}{obj[field.Name], obj[field.Name], obj[field.Name]}
		}
	}
	return obj, nil
}

// ExampleJSON is like Example, but returns a JSON string.
func (d *Definition) ExampleJSON(o Object) (string, error) {
	example, err := d.Example(o)
	if err != nil {
		return "", err
	}
	exampleBytes, err := json.MarshalIndent(example, "", "\t")
	if err != nil {
		return "", err
	}
	return string(exampleBytes), nil
}
