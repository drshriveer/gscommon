package gerrors

import (
	"runtime"
	"strconv"
	"strings"
)

// A Stack represents each line of a Stack trace.
type Stack []StackElem

// String returns a formatted sting with all.
func (s Stack) String() string {
	sb := strings.Builder{}
	for _, e := range s {
		sb.WriteString("\n" + e.Name)
		sb.WriteString("\n\t" + e.File + ":" + strconv.Itoa(e.LineNumber))
	}
	return sb.String()
}

// NearestExternal finds the firs caller outside this package.
// The effectiveness of this method is limited to the depth of the stack fetched.
func (s Stack) NearestExternal() StackElem {
	// attempt to find the first element not in this package:
	if pkgName, ok := getCurrentPackage(); ok {
		for _, elem := range s {
			if !strings.HasPrefix(elem.Name, pkgName) {
				return elem
			}
		}
	}
	return s[0]
}

func getCurrentPackage() (string, bool) {
	pc, _, _, _ := runtime.Caller(1)
	splitName := strings.Split(pcToStackElem(pc).Name, ".")
	if len(splitName[:]) == 0 { // should literally be impossible?
		return "", false
	}
	return strings.Join(splitName[:len(splitName)-1], "."), true
}

// StackElem represents a single line in a Stack trace.
type StackElem struct {
	// Name is the fully qualified package function path.
	// e.g. github.com/drshriveer/gtools/gerrors.TestGError_WithStack
	Name string

	// File is the full path of the file.
	File string

	// LineNumber of the stack element.
	LineNumber int
}

// SourceInfo returns info about where this error was propagated including packageName, typeName, and functionName.
func (e StackElem) SourceInfo() (packageName string, typeName string, funcName string) {
	splitName := strings.Split(e.Name, "/")
	last := splitName[len(splitName)-1]
	// Next step: handle generics which show up as funcName[....]
	// I'd love to do:
	//    last = strings.Replace(last, "[...]", "[T]", 1)
	// but this probably isn't metric safe.
	// I also assume that [...] handles N types so [T] wouldn't quite work.
	last = strings.TrimSuffix(last, "[...]")
	vals := strings.Split(last, ".")

	if len(vals) == 2 {
		packageName, funcName = vals[0], vals[1]
	} else if len(vals) == 3 {
		packageName, typeName, funcName = vals[0], vals[1], vals[2]
		typeName = strings.TrimPrefix(typeName, "(")
		typeName = strings.TrimSuffix(typeName, ")")
	}

	return packageName, typeName, funcName
}

// Metric returns a metric-safe(?) string of the source info.
func (e StackElem) Metric() string {
	pkg, tName, fName := e.SourceInfo()
	tName = strings.TrimPrefix(tName, "*") // remove pointer indicator
	return convertToMetricNode(pkg, tName, fName)
}

func makeStack(depth, skip int) Stack {
	pcs := make([]uintptr, depth)
	n := runtime.Callers(skip, pcs)
	pcs = pcs[0:n] // drop unwritten elements.
	stack := make(Stack, n)
	for i := range stack {
		stack[i] = pcToStackElem(pcs[i])
	}

	return stack
}

func pcToStackElem(pc uintptr) StackElem {
	pc = pc - 1
	fu := runtime.FuncForPC(pc)
	if fu == nil {
		return StackElem{Name: "unknown", File: "unknown"}
	}
	fName, fLine := fu.FileLine(pc)
	fu.Entry()
	return StackElem{Name: fu.Name(), File: fName, LineNumber: fLine}
}
