package oto

import (
	"go/token"
	"go/types"
	"log"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

type parser struct {
	patterns []string
	def      definition
}

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
	pkgs, err := packages.Load(cfg, p.patterns...)
	if err != nil {
		return p.def, err
	}
	for _, pkg := range pkgs {
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
			default:
				log.Println(item, obj)
			}
		}
	}
	return p.def, nil
}

func (p *parser) parseService(pkg *packages.Package, obj types.Object, interfaceType *types.Interface) (service, error) {
	var s service
	s.Name = obj.Name()
	l := interfaceType.NumMethods()
	for i := 0; i < l; i++ {
		m := interfaceType.Method(i)
		method, err := p.parseMethod(pkg, m)
		if err != nil {
			return s, err
		}
		s.Methods = append(s.Methods, method)
	}
	return s, nil
}

func (p *parser) parseMethod(pkg *packages.Package, methodType *types.Func) (method, error) {
	var m method
	m.Name = methodType.Name()
	sig := methodType.Type().(*types.Signature)
	inputParams := sig.Params()
	if inputParams.Len() != 1 {
		return m, p.errWrap(errors.New("service methods must have signature (Request) Response"), pkg, methodType.Pos())
	}
	m.InputObject = p.parseType(pkg, inputParams.At(0))
	outputParams := sig.Results()
	if outputParams.Len() != 1 {
		return m, p.errWrap(errors.New("service methods must have signature (Request) Response"), pkg, methodType.Pos())
	}
	m.OutputObject = p.parseType(pkg, outputParams.At(0))
	return m, nil
}

func (p *parser) parseObject(pkg *packages.Package, o types.Object, v *types.Struct) error {
	var obj object
	obj.Name = o.Name()
	typ := v.Underlying()
	st, ok := typ.(*types.Struct)
	if !ok {
		return p.errWrap(errors.New(obj.Name+" must be a struct"), pkg, o.Pos())
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
		return f, p.errWrap(errors.New(f.Name+" must be exported"), pkg, v.Pos())
	}
	f.Type = p.parseType(pkg, v)
	return f, nil
}

func (p *parser) parseType(pkg *packages.Package, obj types.Object) fieldType {
	resolver := func(other *types.Package) string {
		if other.Name() != pkg.Name {
			if p.def.Imports == nil {
				p.def.Imports = make(map[string]bool)
			}
			p.def.Imports[other.Path()] = true
			return other.Name()
		}
		return ""
	}
	var ftype fieldType
	typ := obj.Type()
	if slice, ok := obj.Type().(*types.Slice); ok {
		typ = slice.Elem()
		ftype.Multiple = true
	}
	ftype.Name = types.TypeString(typ, resolver)
	return ftype
}

func (p *parser) errWrap(err error, pkg *packages.Package, pos token.Pos) error {
	position := pkg.Fset.Position(pos)
	return errors.Wrap(err, position.String())
}

type definition struct {
	Services []service       `json:"services,omitempty"`
	Objects  []object        `json:"objects,omitempty"`
	Imports  map[string]bool `json:"imports,omitempty"`
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
	Name string    `json:"name,omitempty"`
	Type fieldType `json:"type,omitempty"`
}

type fieldType struct {
	Name     string `json:"name,omitempty"`
	Multiple bool   `json:"multiple,omitempty"`
}
