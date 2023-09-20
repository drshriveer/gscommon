package gen

import (
	_ "embed"
	"errors"
	"fmt"
	"go/types"
	"slices"
	"sort"
	"text/template"

	"github.com/fatih/structtag"

	"github.com/drshriveer/gtools/gencommon"
	"github.com/drshriveer/gtools/set"
)

var (
	//go:embed gerror.gotmpl
	rawGerrorTemplate string
	sortTemplate      = template.Must(template.New("gerror").Parse(rawGerrorTemplate))
)

// Generate is the parser and writer of gerrors
// It seems to double as its own 'options' holder.
type Generate struct {
	InFile  string   `alias:"in" env:"GOFILE" usage:"path to input file (defaults to go:generate context)"`
	OutFile string   `alias:"out" usage:"name of output file (defaults to go:generate context filename.gerror.go)"`
	Types   []string `usage:"[required] names of types to generate gerrors for"`

	// derived, (exposed for template use):
	Imports    gencommon.ImportHandler
	PkgName    string
	ErrorDescs ErrorDescs
}

// Parse the input file and drives the attributes above.
func (g *Generate) Parse() error {
	_, pkg, _, err := gencommon.LoadPackages(g.InFile)
	if err != nil {
		return err
	}

	iFact := gencommon.FindInterface(pkg.Imports["github.com/drshriveer/gtools/gerrors"], "Factory")
	methods := iFact.Methods.List
	comments := make(map[string]string, len(methods))
	for _, m := range methods {
		comments[m.Names[0].Name] = gencommon.CommentGroupRaw(m.Doc)
	}

	pkg.Types.Scope()
	g.Imports = gencommon.CalcImports(pkg)
	g.PkgName = pkg.Name
	pkgScope := pkg.Types.Scope()

	g.ErrorDescs = make(ErrorDescs, len(g.Types))
	for i, errType := range g.Types {
		obj := pkgScope.Lookup(errType)
		g.ErrorDescs[i], err = g.createErrorDesc(obj, errType)
		g.ErrorDescs[i].Comments = comments
		if err != nil {
			return err
		}
	}

	return nil
}

// Write writes out the enum config file as configured.
func (g *Generate) Write() error {
	return gencommon.Write(sortTemplate, g, g.OutFile)
}

func (g *Generate) createErrorDesc(obj types.Object, typeName string) (*ErrorDesc, error) {
	if obj == nil {
		return nil, errors.New(typeName + " was not found in AST")
	}

	strukt, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		return nil, errors.New(typeName + " is not a struct")
	}

	fields := make(Fields, 0, strukt.NumFields())
	embedsGError := false
	for i := 0; i < strukt.NumFields(); i++ {
		ff := strukt.Field(i)
		field, err := g.createField(ff, strukt.Tag(i))
		if err != nil {
			return nil, err
		} else if field != nil {
			fields = append(fields, field)
		}

		if ff.Embedded() && ff.Name() == "GError" {
			embedsGError = true
		}
	}

	if !embedsGError {
		return nil, errors.New(typeName + " does not embed GError")
	}

	sort.Sort(fields)
	return &ErrorDesc{
		TypeName: typeName,
		Fields:   fields,
	}, nil
}

func (g *Generate) createField(field *types.Var, tagLine string) (*Field, error) {
	tags, err := structtag.Parse(tagLine)
	if err != nil { // error returned when not found
		return nil, nil
	}

	gerrTags, err := tags.Get("gerror")
	if err != nil { // error returned when not found
		return nil, nil
	}
	validOptions := set.Make("clone", "print")
	if !validOptions.Has(gerrTags.Options...) {
		return nil, fmt.Errorf("field %s has unsupported options; vald=%+v found=%+v",
			field.Name(), validOptions, gerrTags.Options)
	}

	if gerrTags.Name == "_" {
		gerrTags.Name = field.Name()
	}
	_ = g.Imports.ExtractTypeRef(field.Type()) // just register the use of an imported type.
	return &Field{
		Name:    field.Name(),
		PrintAs: gerrTags.Name,
		Clone:   slices.Contains(gerrTags.Options, "clone"),
		Print:   slices.Contains(gerrTags.Options, "print"),
	}, nil
}
