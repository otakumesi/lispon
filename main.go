package main

import (
	"fmt"

	"./lisp"
)

func main() {
	lisp.Init()
	sym := lisp.Symbol{Name: lisp.String("+"), IsQuoted: false}
	cons := lisp.Cons{lisp.Number(1), lisp.Cons{lisp.Number(5), lisp.Nil{}}}
	sexpr := lisp.NewSExpr(sym, cons)
	fmt.Println(lisp.Eval(sexpr))
}
