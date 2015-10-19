package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

var fmtPrintf = fmt.Printf
var fileNamePtr *string

func init() {
	fileNamePtr = flag.String("f", "", "Filename to parse")
}

func main() {
	file := getFileFromArgs()
	fields := getFieldsFromFile(file)
	//TODO: output variables for highlighting
	//TODO: method parameters from ast
	fmtPrintf(strings.Join(fields, " "))
}

func getFileFromArgs() string {
	flag.Parse()
	return *fileNamePtr
}

func getFieldsFromFile(file string) []string {
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var foundFields []string
	for _, declaration := range f.Decls {
		if genDecl, isGenDecl := declaration.(*ast.GenDecl); isGenDecl {
			// fmt.Printf("genDecl: %q isGenDecl: %q\n", genDecl, isGenDecl)
			for _, s := range genDecl.Specs {
				if value, isValueSpec := s.(*ast.ValueSpec); isValueSpec {
					// fmt.Printf("\tvariable: %q\n", value.Names[0])

					foundFields = append(foundFields, value.Names[0].String())
				}
			}
		}
	}
	return foundFields
}

type method struct {
	bodyStart int      //Line the method starts at
	bodyEnd   int      //Line the method ends at
	variables []string //a list of variables
}

func getVarsFromFile(file string) method {
	fset := token.NewFileSet() // positions are relative to fset
	m := method{}

	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		fmt.Println(err)
		return m
	}

	for _, declaration := range f.Decls {
		//Function declarations
		if funcDecl, isFuncDecl := declaration.(*ast.FuncDecl); isFuncDecl {
			// fmt.Printf("fu: %q\n", funcDecl)

			startingPostition := fset.Position(funcDecl.Pos())
			m.bodyStart = startingPostition.Line

			endingPostition := fset.Position(funcDecl.End())
			m.bodyEnd = endingPostition.Line

			for _, methodStatement := range funcDecl.Body.List {
				// fmt.Printf("methodStatement: %q\n", methodStatement)
				if assignStmt, isAssignment := methodStatement.(*ast.AssignStmt); isAssignment {
					for _, itemLeftOfAssignment := range assignStmt.Lhs {
						if variable, isVarIdent := itemLeftOfAssignment.(*ast.Ident); isVarIdent {
							m.variables = append(m.variables, variable.Name)
						}
					}
				}
			}
		}
	}
	return m
}
