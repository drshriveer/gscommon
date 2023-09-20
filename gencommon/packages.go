package gencommon

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path"
	"path/filepath"
)

func LoadPackages(inFile string) (*token.FileSet, *packages.Package, *ast.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedDeps | packages.NeedTypesInfo | packages.NeedTypes |
			packages.NeedEmbedPatterns | packages.NeedSyntax,
	}
	pkgs, err := packages.Load(cfg, path.Dir(inFile))
	if err != nil {
		return nil, nil, nil, err
	}
	if len(pkgs) < 1 {
		return nil, nil, nil, errors.New("package " + inFile + " NOT FOUND")
	}

	fs := token.NewFileSet()
	dir := dir(pkgs[0])
	astPkgs, err := parser.ParseDir(fs, dir, nil, parser.DeclarationErrors|parser.ParseComments)
	if err != nil {
		return nil, nil, nil, err
	}

	if ap, ok := astPkgs[pkgs[0].Name]; ok {
		return fs, pkgs[0], ap, nil
	}

	return fs, pkgs[0], &ast.Package{Name: pkgs[0].Name}, nil
}

func dir(p *packages.Package) string {
	files := append(p.GoFiles, p.OtherFiles...)
	if len(files) < 1 {
		return p.PkgPath
	}

	return filepath.Dir(files[0])
}
