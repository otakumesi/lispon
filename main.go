package main

import (
	"bufio"
	"fmt"
	"os"

	"./lisp"
	"./parser"
)

func main() {
	repl()
}

func repl() {
	lisp.Init()
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for stdin.Scan() {
		text := stdin.Text()
		rootAst := parser.Parse(text)
		sexpr := parser.ParseSExpr(rootAst)
		fmt.Println("=> ",  lisp.Eval(sexpr))
		fmt.Print("> ")
	}
}
