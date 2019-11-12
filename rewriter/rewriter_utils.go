package rewriter

import (
	"go/ast"
)

// RewriteStructArguments handles a case like this
//
// ```go
// func before(a struct { b, c, d int })
//
// func after(a *struct { b, c, d int })
// ```
func RewriteStructArguments(n ast.Node) (ast.Node, bool) {
	_, ok := n.(*ast.FuncDecl)
	if !ok {
		return n, true
	}

	return n, true
}

func isRewriteableFunc(fd *ast.FuncDecl) bool {
	return true
}
