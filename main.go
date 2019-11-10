package main

import (
	"bufio"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"

	"github.com/ivzhh/go-ast-transform/rewriter"

	"github.com/fatih/astrewrite"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	pkgname := reflect.TypeOf(struct{}{}).PkgPath()

	tmpdir := filepath.Join(os.TempDir(), pkgname)

	if err := os.MkdirAll(tmpdir, 0644); err != nil {
		log.Fatalf("create output dir fail: %s", err.Error())
	}

	inputfile := path.Join("testdata", "example1.go")

	args := os.Args[1:]

	if len(args) > 1 {
		log.Fatalf("please run command: %s path_to_go_file.go", pkgname)
	} else if len(args) == 1 {
		inputfile = args[0]
	}

	if _, err := os.Stat(inputfile); os.IsNotExist(err) {
		log.Fatalf("input go file doesn't exist: %s", inputfile)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, inputfile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	rewritten := astrewrite.Walk(file, rewriter.RewriteStructArguments)

	// var buf bytes.Buffer
	// printer.Fprint(&buf, fset, rewritten)
	// fmt.Println(buf.String())

	var f *os.File

	outputfile := filepath.Join(tmpdir, filepath.Base(inputfile))

	if f, err = os.Create(outputfile); err != nil {
		log.Fatalf("create file: %+v", err)
	}

	log.Printf("output to %s", outputfile)

	ast.Fprint(bufio.NewWriter(f), fset, rewritten, nil)
}
