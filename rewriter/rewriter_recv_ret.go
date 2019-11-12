package rewriter

import (
	"go/ast"
	"go/token"
)

// RewriteReturnVars handles a case like this
//
// ```go
// func before(a struct { b, c, d int }) (r int)
//
// func after(a struct { b, c, d int }, r *int)
// ```
func RewriteReturnVars(n ast.Node) (ast.Node, bool) {
	_, ok := n.(*ast.FuncDecl)
	if !ok {
		return n, true
	}

	return n, true
}

type Rewriter struct {
	Original *ast.File
	Fset     *token.FileSet
}
