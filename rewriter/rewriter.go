package rewriter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

// Rewriter provides a shared context for all rewriting
type Rewriter struct {
	Fset  *token.FileSet
	Files map[string]*FileRewriter
}

// NewRewritter creates a new Rewriter object
// if `fset` is `nil`, then create a new one
func NewRewritter(fset *token.FileSet) *Rewriter {
	if fset == nil {
		fset = token.NewFileSet()
	}

	return &Rewriter{
		Fset:  fset,
		Files: map[string]*FileRewriter{},
	}
}

// NewFileRewritter creates a FileRewriter for specific file
func (r *Rewriter) NewFileRewritter(inputfile string) *FileRewriter {
	var err error
	var file *ast.File
	if file, err = parser.ParseFile(r.Fset, inputfile, nil, parser.ParseComments); err != nil {
		log.Fatal(err)
		return nil
	}

	var checker types.Config = types.Config{}

	info := types.Info{}

	if _, err = checker.Check(inputfile, r.Fset, []*ast.File{file}, &info); err != nil {
		log.Fatal(err)
		return nil
	}

	return &FileRewriter{
		ctx:      r,
		Original: file,
		Wrapper:  nil,
	}
}
