package rewriter

import (
	"go/ast"
	"log"
)

// Rewriter modifies ast node and returns a new one
type Rewriter func(n ast.Node) (ast.Node, bool)

// RewriteReturnVars handles a case like this
//
// ```go
// func before(a struct { b, c, d int }) (r int)
//
// func after(a *struct { b, c, d int }, r *int)
// ```
func RewriteReturnVars(n ast.Node) (ast.Node, bool) {
	fd, ok := n.(*ast.FuncDecl)
	if !ok {
		return n, true
	}

	log.Printf("%+v", fd.Name.Obj.Decl)

	return n, true
}

// RewriteStructArguments handles a case like this
//
// ```go
// func before(a struct { b, c, d int })
//
// func after(a *struct { b, c, d int })
// ```
func RewriteStructArguments(n ast.Node) (ast.Node, bool) {
	fd, ok := n.(*ast.FuncDecl)
	if !ok {
		return n, true
	}

	log.Printf("%+v", fd.Name.Obj.Decl)

	return n, true
}
