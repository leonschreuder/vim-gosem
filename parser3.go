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
var fset *token.FileSet

func init() {
	fileNamePtr = flag.String("f", "", "Filename to parse")
}

//TODO: method parameters from ast

func main() {
	file := getFileFromArgs()

	var groups []string

	fields := getFieldsFromFile(file)
	groups = append(groups, strings.Join(fields, " "))

	methods := getVarsFromFile(file)
	for _, currentMethod := range methods {
		varString := strings.Join(currentMethod.variables, " ")
		methodString := fmt.Sprintf("%d,%d,%s", currentMethod.bodyStart, currentMethod.bodyEnd, varString)

		groups = append(groups, methodString)
	}

	fmtPrintf(strings.Join(groups, "|"))
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

func getVarsFromFile(file string) []method {
	fset = token.NewFileSet() // positions are relative to fset
	var ms []method

	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		fmt.Println(err)
		return ms
	}

	for _, declaration := range f.Decls {
		if funcDecl, isFuncDecl := declaration.(*ast.FuncDecl); isFuncDecl {
			m := traverseFunction(fset, funcDecl)
			ms = append(ms, m)
		}
	}
	return ms
}

func traverseFunction(fset *token.FileSet, funcDecl *ast.FuncDecl) method {
	m := method{}

	// ast.Print(fset, methodStatement)

	startingPostition := fset.Position(funcDecl.Pos())
	m.bodyStart = startingPostition.Line

	endingPostition := fset.Position(funcDecl.End())
	m.bodyEnd = endingPostition.Line
	m.variables = getVariablesFromMethod(funcDecl.Body)
	return m
}

func getVariablesFromMethod(methodBlock *ast.BlockStmt) []string {
	var allVars []string
	for _, methodStatement := range methodBlock.List {
		if decl, isDecl := methodStatement.(*ast.DeclStmt); isDecl {
			allVars = append(allVars, getVariablesFromDeclStmt(decl)...)
		} else if assignStmt, isAssignment := methodStatement.(*ast.AssignStmt); isAssignment {
			allVars = append(allVars, getVariablesFromAssignStmt(assignStmt)...)
		}
	}
	return allVars
}

func getVariablesFromAssignStmt(assignStmt *ast.AssignStmt) []string {
	var vars []string
	for _, itemLeftOfAssignment := range assignStmt.Lhs {
		if variable, isVarIdent := itemLeftOfAssignment.(*ast.Ident); isVarIdent {
			vars = append(vars, variable.Name)
		}
	}
	return vars
}

func getVariablesFromDeclStmt(decl *ast.DeclStmt) []string {
	var vars []string
	if genDecl, isGenDecl := decl.Decl.(*ast.GenDecl); isGenDecl {
		for _, s := range genDecl.Specs {
			if value, isValueSpec := s.(*ast.ValueSpec); isValueSpec {
				for _, ident := range value.Names {
					vars = append(vars, ident.Name)
				}
			}
		}
	}
	return vars
}
