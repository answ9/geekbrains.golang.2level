package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

func main() {
	count, err := CountAsyncFunctions("fileWithCode.go", "TheFunctionWithSomeAsyncFunctionsInside")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Async functions found:", count)
}

func CountAsyncFunctions(fileName, funcName string) (count int, err error) {
	functionExists := false
	set := token.NewFileSet()
	astFile, err := parser.ParseFile(set, fileName, nil, 0)
	if err != nil {
		return
	}

	for _, d := range astFile.Decls {
		if fn, isFunc := d.(*ast.FuncDecl); isFunc {
			if fn.Name.String() != funcName {
				continue
			}
			functionExists = true
			count = countGoStmt(fn.Body.List)
			break
		}
	}
	if !functionExists {
		err = fmt.Errorf("function with name \"%s\" was not found", funcName)
	}
	return
}

func countGoStmt(stmts []ast.Stmt) int {
	var count int
	for _, stmt := range stmts {
		reflect.ValueOf(stmt)
		switch v := stmt.(type) {
		case *ast.GoStmt:
			count++
		case *ast.IfStmt:
			count += countGoStmt(v.Body.List)
		case *ast.ForStmt:
			count += countGoStmt(v.Body.List)
		case *ast.RangeStmt:
			count += countGoStmt(v.Body.List)
		case *ast.SwitchStmt:
			count += countGoStmt(v.Body.List)
		case *ast.CaseClause:
			count += countGoStmt(v.Body)
		}
	}
	return count
}
