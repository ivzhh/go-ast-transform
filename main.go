package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"

	"github.com/fatih/astrewrite"
	"github.com/ivzhh/go-ast-transform/rewriter"
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

	var err error

	r := rewriter.NewRewritter(fset)
	fr := r.NewFileRewritter(inputfile)

	rewritten := astrewrite.Walk(fr.Original, fr.RewriteReturnVars())

	{
		var f *os.File

		fmtfile := filepath.Join(tmpdir, fmt.Sprintf("fmt-%s", filepath.Base(inputfile)))

		if f, err = os.Create(fmtfile); err != nil {
			log.Fatalf("create file: %+v", err)
		}

		printer.Fprint(f, fset, rewritten)

		log.Printf("output to %s", fmtfile)

		defer func() { f.Close() }()
	}
	{
		var f *os.File

		outputfile := filepath.Join(tmpdir, fmt.Sprintf("ast-%s", filepath.Base(inputfile)))

		if f, err = os.Create(outputfile); err != nil {
			log.Fatalf("create file: %+v", err)
		}

		log.Printf("output to %s", outputfile)

		if err = ast.Fprint(f, fset, rewritten, nil); err != nil {
			log.Fatalf("output ast to file: %+v", err)
		}

		defer func() { f.Close() }()
	}
}
