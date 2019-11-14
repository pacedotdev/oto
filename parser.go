package main

import (
	"fmt"
	"go/token"
	"go/types"
	"sort"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

var errNotFound = errors.New("not found")

type definition struct {
	PackageName string            `json:"packageName,omitempty"`
	Services    []service         `json:"services,omitempty"`
	Objects     []object          `json:"objects,omitempty"`
	Imports     map[string]string `json:"imports,omitempty"`
}

// Object looks up an object by name. Returns errNotFound error
// if it cannot find it.
func (d *definition) Object(name string) (*object, error) {
	for i := range d.Objects {
		obj := &d.Objects[i]
		if obj.Name == name {
			return obj, nil
		}
	}
	return nil, errNotFound
}

type service struct {
	Name    string   `json:"name,omitempty"`
	Methods []method `json:"methods,omitempty"`
}

type method struct {
	Name         string    `json:"name,omitempty"`
	InputObject  fieldType `json:"inputObject,omitempty"`
	OutputObject fieldType `json:"outputObject,omitempty"`
}

type object struct {
	Name   string  `json:"name,omitempty"`
	Fields []field `json:"fields,omitempty"`
}

type field struct {
	Name      string    `json:"name,omitempty"`
	Type      fieldType `json:"type,omitempty"`
	OmitEmpty bool      `json:"omitEmpty,omitempty"`
}

type fieldType struct {
	TypeName string `json:"typeName,omitempty"`
	Multiple bool   `json:"multiple,omitempty"`
	Package  string `json:"package,omitempty"`
}

type parser struct {
	patterns      []string
	def           definition
	outputObjects map[string]bool
	Verbose       bool
}

// newParser makes a fresh parser using the specified patterns.
// The patterns should be the args passed into the tool (after any flags)
// and will be passed to the underlying build system.
func newParser(patterns ...string) *parser {
	return &parser{
		patterns: patterns,
	}
}

func (p *parser) parse() (definition, error) {
	cfg := &packages.Config{
		Mode:  packages.NeedTypes | packages.NeedDeps | packages.NeedName,
		Tests: false,
	}
	p.outputObjects = make(map[string]bool)
	pkgs, err := packages.Load(cfg, p.patterns...)
	if err != nil {
		return p.def, err
	}
	for _, pkg := range pkgs {
		p.def.PackageName = pkg.Name
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			switch item := obj.Type().Underlying().(type) {
			case *types.Interface:
				s, err := p.parseService(pkg, obj, item)
				if err != nil {
					return p.def, err
				}
				p.def.Services = append(p.def.Services, s)
			case *types.Struct:
				p.parseObject(pkg, obj, item)
			}
		}
	}
	sort.Slice(p.def.Services, func(i, j int) bool {
		return p.def.Services[i].Name < p.def.Services[j].Name
	})
	if err := p.addOutputFields(); err != nil {
		return p.def, err
	}
	return p.def, nil
}

func (p *parser) parseService(pkg *packages.Package, obj types.Object, interfaceType *types.Interface) (service, error) {
	var s service
	s.Name = obj.Name()
	if p.Verbose {
		fmt.Printf("%s ", s.Name)
	}
	l := interfaceType.NumMethods()
	for i := 0; i < l; i++ {
		m := interfaceType.Method(i)
		method, err := p.parseMethod(pkg, s.Name, m)
		if err != nil {
			return s, err
		}
		s.Methods = append(s.Methods, method)
	}
	return s, nil
}

func (p *parser) parseMethod(pkg *packages.Package, serviceName string, methodType *types.Func) (method, error) {
	var m method
	m.Name = methodType.Name()
	sig := methodType.Type().(*types.Signature)
	inputParams := sig.Params()
	if inputParams.Len() != 1 {
		return m, p.wrapErr(errors.New("invalid method signature: expected Method(MethodRequest) MethodResponse"), pkg, methodType.Pos())
	}
	m.InputObject = p.parseType(pkg, inputParams.At(0))
	outputParams := sig.Results()
	if outputParams.Len() != 1 {
		return m, p.wrapErr(errors.New("invalid method signature: expected Method(MethodRequest) MethodResponse"), pkg, methodType.Pos())
	}
	m.OutputObject = p.parseType(pkg, outputParams.At(0))
	p.outputObjects[m.OutputObject.TypeName] = true
	return m, nil
}

// parseObject parses a struct type and adds it to the definition.
func (p *parser) parseObject(pkg *packages.Package, o types.Object, v *types.Struct) error {
	var obj object
	obj.Name = o.Name()
	typ := v.Underlying()
	st, ok := typ.(*types.Struct)
	if !ok {
		return p.wrapErr(errors.New(obj.Name+" must be a struct"), pkg, o.Pos())
	}
	for i := 0; i < st.NumFields(); i++ {
		field, err := p.parseField(pkg, st.Field(i))
		if err != nil {
			return err
		}
		obj.Fields = append(obj.Fields, field)
	}
	p.def.Objects = append(p.def.Objects, obj)
	return nil
}

func (p *parser) parseField(pkg *packages.Package, v *types.Var) (field, error) {
	var f field
	f.Name = v.Name()
	if !v.Exported() {
		return f, p.wrapErr(errors.New(f.Name+" must be exported"), pkg, v.Pos())
	}
	f.Type = p.parseType(pkg, v)
	return f, nil
}

func (p *parser) parseType(pkg *packages.Package, obj types.Object) fieldType {
	var ftype fieldType
	resolver := func(other *types.Package) string {
		if other.Name() != pkg.Name {
			if p.def.Imports == nil {
				p.def.Imports = make(map[string]string)
			}
			p.def.Imports[other.Path()] = other.Name()
			ftype.Package = other.Path()
			return other.Name()
		}
		return "" // no package prefix
	}
	typ := obj.Type()
	if slice, ok := obj.Type().(*types.Slice); ok {
		typ = slice.Elem()
		ftype.Multiple = true
	}
	ftype.TypeName = types.TypeString(typ, resolver)
	return ftype
}

// addOutputFields adds built-in fields to the response objects
// mentioned in p.outputObjects.
func (p *parser) addOutputFields() error {
	errorField := field{
		OmitEmpty: true,
		Name:      "Error",
		Type: fieldType{
			TypeName: "string",
		},
	}
	for typeName := range p.outputObjects {
		obj, err := p.def.Object(typeName)
		if err != nil {
			return errors.Wrapf(err, "missing output object: %s", typeName)
		}
		obj.Fields = append(obj.Fields, errorField)
	}
	return nil
}

func (p *parser) wrapErr(err error, pkg *packages.Package, pos token.Pos) error {
	position := pkg.Fset.Position(pos)
	return errors.Wrap(err, position.String())
}
