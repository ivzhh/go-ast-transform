package rewriter

import (
	"go/ast"
	"log"
)

// FileRewriter is the context of rewriting process
type FileRewriter struct {
	ctx      *Rewriter
	Original *ast.File

	Wrapper *ast.File
}

// RewriteReturnVars handles a case like this
//
// ```go
// func before(a struct { b, c, d int }) (r int)
//
// func after(a struct { b, c, d int }, r *int)
// ```
func (config *FileRewriter) RewriteReturnVars() func(n ast.Node) (ast.Node, bool) {
	return func(n ast.Node) (ast.Node, bool) {
		fd, ok := n.(*ast.FuncDecl)
		if !ok {
			return n, true
		}

		hasRewriteableRecv := false
		if fd.Recv != nil || fd.Recv.NumFields() > 0 {
			if len(fd.Recv.List) != 1 {
				return n, false
			}

			recv := fd.Recv.List[0]

			switch typ := recv.Type.(type) {
			case *ast.StarExpr:
			case *ast.Ident:
				isPointer, err := traceIsPointer(typ)

				if err != nil {
					log.Printf("fail to parse star expression")
					return n, false
				}

				if !isPointer {
					recv.Type = &ast.StarExpr{
						X: typ,
					}
				}

			}
		}

		if hasRewriteableRecv || fd.Type.Results != nil || fd.Type.Results.NumFields() > 0 {
			config.ensureWrapperFile()
		} else {
			return n, true
		}

		return n, true
	}
}

func (config *FileRewriter) ensureWrapperFile() {
	if config.Wrapper == nil {
		config.Wrapper = &ast.File{
			Doc:        &ast.CommentGroup{},
			Package:    config.Original.Package,
			Name:       config.Original.Name,
			Decls:      []ast.Decl{},
			Scope:      ast.NewScope(nil),
			Imports:    []*ast.ImportSpec{},
			Unresolved: []*ast.Ident{},
			Comments:   []*ast.CommentGroup{},
		}
	}
}
