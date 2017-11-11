package main

import (
	"fmt"

	"./lisp"
)

func main() {

	sym := lisp.Symbol("+")
	lhs := lisp.Number(1)
	cons_l := lisp.Number(2)
	cons_r := lisp.Cons{lisp.Number(5), lisp.Nil{}}
	rhs := lisp.Cons{cons_l, cons_r}
	sexpr := lisp.NewSExpr(sym, lhs, rhs)
	fmt.Println(sexpr.Eval())

}
