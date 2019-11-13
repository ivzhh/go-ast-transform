package rewriter

import (
	"go/ast"
	"log"
	"reflect"
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

func traceTypedef(ident *ast.Ident) {
	var definedFrom ast.Expr

	switch spec := ident.Obj.Decl.(type) {
	case *ast.TypeSpec:
		definedFrom = spec.Type
		log.Printf("%+v", reflect.TypeOf(definedFrom))
	default:
		log.Fatalf("expected *ast.TypeSpec, but get %+v", reflect.TypeOf(spec))
	}

	switch definedFrom.(type) {
	case *ast.StructType:
	case *ast.StarExpr:
	}
}
