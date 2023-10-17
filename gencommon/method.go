package gencommon

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
)

// Method is a type representing a method.
type Method struct {
	Name       string `gsort:"2"`
	Comments   Comments
	Input      Params
	Output     Params
	IsExported bool `gsort:"1"`
}

// MethodFromSignature returns a partially constructed Method-- it contains the raw function
// info, nothing else.
func MethodFromSignature(ih *ImportHandler, signature *types.Signature) *Method {
	return &Method{
		Name:   "func", // default name
		Input:  ParamsFromSignatureTuple(ih, signature.Params()),
		Output: ParamsFromSignatureTuple(ih, signature.Results()),
	}
}

// Signature returns the full signature of the method.
// e.g. MethodName(arg1 Type1, arg2 Type2, arg3 Type3) (*Thing, string, error).
func (m *Method) Signature() string {
	if len(m.Output) > 1 {
		return fmt.Sprintf("%s(%s) (%s)", m.Name, m.Input.Declarations(), m.Output.TypeNames())
	}
	return fmt.Sprintf("%s(%s) %s", m.Name, m.Input.Declarations(), m.Output.TypeNames())
}

// Call returns a method call of the form:
// e.g. MethodName(arg1, arg2, arg3).
func (m *Method) Call() string {
	args := mapper(m.Input, func(in *Param) string { return in.Name })
	return fmt.Sprintf("%s(%s) ", m.Name, strings.Join(args, ", "))
}

// ReturnsError indicates whether this method return a type that implements the error
// interface as its last argument.
func (m *Method) ReturnsError() bool {
	if len(m.Output) == 0 {
		return false
	}
	last := m.Output[len(m.Output)-1]
	return types.Implements(last.actualType, ErrorInterface)
}

func getName(names ...*ast.Ident) string {
	for _, name := range names {
		return name.Name
	}

	return ""
}

// methodIdent returns true if function declaration is a method of the target type.
func methodIdent(decl *ast.FuncDecl, targetType string) bool {
	if decl.Recv == nil || len(decl.Recv.List) != 1 {
		return false
	}

	switch t := decl.Recv.List[0].Type.(type) {
	case *ast.StarExpr:
		ident, ok := t.X.(*ast.Ident)
		if !ok {
			return false
		}
		return ident.Name == targetType
	case *ast.Ident:
		return t.Name == targetType
	}
	return false
}
