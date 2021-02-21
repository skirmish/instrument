package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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
		var fileSet *token.FileSet = token.NewFileSet()
		inAst, err := parser.ParseFile(fileSet, in, src, parser.ParseComments)

		// log some info
		t.Log(inAst.Name.Name)
		ast.Print(fileSet, inAst)

		// now turn the ast back into source
		var buf bytes.Buffer
		printer.Fprint(&buf, fileSet, inAst)

		// and log
		t.Log("Output file:\n", buf.String())
	}
}
