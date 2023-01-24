package parser

import "fmt"

// Example generates an object that is a realistic example
// of this object.
// Examples are read from the docs.
// This is experimental.
func (d *Definition) Example(o Object) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	for _, field := range o.Fields {
		if field.Type.IsObject {
			subobj, err := d.Object(field.Type.CleanObjectName)
			if err != nil {
				return nil, fmt.Errorf("Object(%q): %w", field.Type.CleanObjectName, err)
			}
			example, err := d.ExampleP(subobj)
			if err != nil {
				return nil, err
			}
			obj[field.NameLowerCamel] = example
			continue
		}
		obj[field.NameLowerCamel] = field.Example
	}
	return obj, nil
}

// ExampleP is a pointer version of Example.
func (d *Definition) ExampleP(o *Object) (map[string]interface{}, error) {
	return d.Example(*o)
}
