package main

import (
	"bufio"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"

	"github.com/ivzhh/go-ast-transform/rewriter"

	"github.com/fatih/astrewrite"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path.Join("testdata", "example1.go"), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	rewritten := astrewrite.Walk(file, rewriter.RewriteStructArguments)

	// var buf bytes.Buffer
	// printer.Fprint(&buf, fset, rewritten)
	// fmt.Println(buf.String())

	var f *os.File

	if f, err = os.Create("/tmp/dat2"); err != nil {
		log.Fatalf("create file: %+v", err)
	}

	ast.Fprint(bufio.NewWriter(f), fset, rewritten, nil)
}
