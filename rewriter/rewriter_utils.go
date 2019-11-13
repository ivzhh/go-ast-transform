package rewriter

import (
	"fmt"
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

func traceIsPointer(ident *ast.Ident) (bool, error) {
	var definedFrom ast.Expr

	switch spec := ident.Obj.Decl.(type) {
	case *ast.TypeSpec:
		definedFrom = spec.Type
	default:
		log.Fatalf("expected *ast.TypeSpec, but get %+v", reflect.TypeOf(spec))
	}

	switch identFrom := definedFrom.(type) {
	case *ast.StructType:
		return false, nil
	case *ast.StarExpr:
		return true, nil
	case *ast.Ident:
		return traceIsPointer(identFrom)
	}

	return false, fmt.Errorf("typedef can only be *ast.StructType, *ast.StarExpr, *ast.Ident")
}
