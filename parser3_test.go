package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	// "leonmoll.de/testutils"
	"os"
	"testing"
)

var originalArgs []string
var mockOut []string // output is written to this var (reset in every func)

func setup() {
}

func teardown() {
}

func Test_main(t *testing.T) {
	addToArgs("-f=input.go")
	defer resetArgs()
	initMockPrinter()
	expected := "someFieldToPrint someOtherField|8,16,stringToPrint a b|18,21,a"

	main()

	if mockOut[0] != expected {
		t.Errorf("\nExpected: %q\nGot     : %q", expected, mockOut[0])
	}
}

func Test_parsingAStringFlag(t *testing.T) {
	inputString := "input.go"
	addToArgs("-f=" + inputString)
	defer resetArgs()

	stringResult := getFileFromArgs()

	if stringResult != inputString {
		t.Errorf("Should parsed a flag with value %q, but got %q", inputString, stringResult)
	}
}

func Test_gettingVariablesFromMethod(t *testing.T) {
	source := `
package main
func main() {
	var typedVariable string
	compilerTypedVar := "printed"
}`
	expectedVars := []string{"typedVariable", "compilerTypedVar"}

	// ast.Print(fset, fun.Body)
	funStmt, _ := parseSource(source).Decls[0].(*ast.FuncDecl)

	returnedVars := getVariablesFromMethod(funStmt.Body)

	if len(returnedVars) != len(expectedVars) {
		t.Errorf("Fields expected %d, got %d ", len(expectedVars), len(returnedVars))
	}
	for i, cField := range expectedVars {
		if cField != returnedVars[i] {
			t.Errorf("Expected field %q, got %q ", cField, returnedVars[i])
		}
	}
}

func parseSource(src string) *ast.File {
	fset = token.NewFileSet() // positions are relative to fset
	f, _ := parser.ParseFile(fset, "", src, 0)
	return f
}

func Test_gettingVariableName(t *testing.T) {
	expectedFields := []string{
		"someFieldToPrint",
		"someOtherField",
	}

	fieldsFound := getFieldsFromFile("input.go")

	if len(fieldsFound) != len(expectedFields) {
		t.Errorf("Fields expected %d, got %d ", len(expectedFields), len(fieldsFound))
	}
	for i, cField := range expectedFields {
		if cField != fieldsFound[i] {
			t.Errorf("Expected field %q, got %q ", cField, fieldsFound[i])
		}
	}
}

func Test_getMethodVarsFromFile(t *testing.T) {
	expectedMethods := []method{
		method{
			bodyStart: 8,
			bodyEnd:   16,
			variables: []string{"stringToPrint", "a", "b"},
		},
		method{
			bodyStart: 18,
			bodyEnd:   21,
			variables: []string{"a"},
		},
	}

	methodsFound := getVarsFromFile("input.go")

	for i, cMethod := range methodsFound {
		if cMethod.bodyStart != expectedMethods[i].bodyStart {
			t.Errorf("Expected method to start at line  %d, but got line %d", expectedMethods[i].bodyStart, cMethod.bodyStart)
		}
		if cMethod.bodyEnd != expectedMethods[i].bodyEnd {
			t.Errorf("Expected method to end at line  %d, but got line %d", expectedMethods[i].bodyEnd, cMethod.bodyEnd)
		}
		if len(cMethod.variables) != len(expectedMethods[i].variables) {
			t.Fatalf("Expected %d items, got %d. Items: %q", len(expectedMethods[i].variables), len(cMethod.variables), cMethod.variables)
		}
		for j, cVar := range cMethod.variables {
			if cVar != expectedMethods[i].variables[j] {
				t.Errorf("Expected method to have var %q, but had %q", expectedMethods[i].variables[j], cVar)
			}
		}
	}
}

func Test_multipleVariableTypes(t *testing.T) {
	expectedMethods := []method{
		method{
			bodyStart: 5,
			bodyEnd:   10,
			variables: []string{"typedVariable", "compilerTypedVar"},
		},
	}

	methodsFound := getVarsFromFile("input2.go")

	for i, cMethod := range methodsFound {
		if cMethod.bodyStart != expectedMethods[i].bodyStart {
			t.Errorf("Expected method to start at line  %d, but got line %d", expectedMethods[i].bodyStart, cMethod.bodyStart)
		}
		if cMethod.bodyEnd != expectedMethods[i].bodyEnd {
			t.Errorf("Expected method to end at line  %d, but got line %d", expectedMethods[i].bodyEnd, cMethod.bodyEnd)
		}
		if len(cMethod.variables) != len(expectedMethods[i].variables) {
			t.Fatalf("Expected %d items, got %d. Items: %q", len(expectedMethods[i].variables), len(cMethod.variables), cMethod.variables)
		}
		for j, cVar := range cMethod.variables {
			if cVar != expectedMethods[i].variables[j] {
				t.Errorf("Expected method to have var %q, but had %q", expectedMethods[i].variables[j], cVar)
			}
		}
	}
}

// Helpers
//--------------------------------------------------------------------------------
func mockPrintf(unformattedString string, a ...interface{}) (n int, err error) {
	mockOut = append(mockOut, fmt.Sprintf(unformattedString, a...))
	return
}

func initMockPrinter() {
	fmtPrintf = mockPrintf
	mockOut = []string{}
}

func addToArgs(args ...string) {
	originalArgs = os.Args
	os.Args = append(os.Args, args...)
}
func resetArgs() {
	os.Args = originalArgs
}
