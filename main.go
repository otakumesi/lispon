package main

import (
	"fmt"
	"os"

	"./lisp"
	"./parser"
)

func main() {
	RunLisp()
}

func RunLisp() {
	lisp.Init()
	rootAst := parser.Parse(os.Args[1])
	sexpr := parser.ParseSExpr(rootAst)
	fmt.Println(lisp.Eval(sexpr))
}
