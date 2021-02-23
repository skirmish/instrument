package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var fileSet *token.FileSet = token.NewFileSet()

func astPreFunc(cr *astutil.Cursor) bool {
	fmt.Println("pre: ", cr.Name()) //, reflect.TypeOf(cr.Node()).Name())
	var buf bytes.Buffer
	// can we dump this? yes, we can
	err := format.Node(&buf, fileSet, cr.Node())
	if err == nil {
		fmt.Println("Output file of node:\n", buf.String())
	}
	return true
}

func astPostFunc(cr *astutil.Cursor) bool {
	fmt.Println("post: ", cr.Name())
	return true
}

// TestInstrumentation processes files in testdata/*.input
// and compares them against the corresponding testdata/*.expected files.
func TestInstrumentation1(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// grab a list of the input files.
	match, err := filepath.Glob(filepath.Join(path, "/testdata/*.input"))
	if err != nil {
		t.Fatal(err)
	}

	for _, in := range match {
		out := ""
		if strings.HasSuffix(in, ".input") {
			out = in[:len(in)-len(".input")] + ".expected"
		}
		t.Log("Input: ", in)
		t.Log("Output: ", out)

		// read in file.input
		f, err := os.Open(in)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		src, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}

		// ast it

		inAst, err := parser.ParseFile(fileSet, in, src, parser.ParseComments)

		// log some info
		t.Log(inAst.Name.Name)
		ast.Print(fileSet, inAst)

		// so now we have an AST, we can walk it
		outAst := astutil.Apply(inAst, astPreFunc, nil)

		// now turn the ast back into source
		var buf bytes.Buffer
		//printer.Fprint(&buf, fileSet, inAst)
		err = format.Node(&buf, fileSet, outAst)

		// and log
		t.Log("Output file:\n", buf.String())
	}
}
