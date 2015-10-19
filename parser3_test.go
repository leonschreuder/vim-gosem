package main

import (
	"fmt"
	"os"
	"testing"
)

var originalArgs []string
var mockOut []string // output is written to this var (reset in every func)

func Test_main(t *testing.T) {
	addToArgs("-f=input.go")
	defer resetArgs()
	expectedOut := "someFieldToPrint someOtherField|8,16,stringToPrint a b|18,21,a"
	initMockPrinter()

	main()

	if mockOut[0] != expectedOut {
		t.Errorf("Expected %q, got %q", expectedOut, mockOut[0])
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
