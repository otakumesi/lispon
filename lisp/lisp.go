package lisp

import "fmt"

type Evaluable interface {
	eval(...Scope) Evaluable
}

func Eval(l Evaluable, scs ...Scope) Evaluable {
	return l.eval(scs...)
}

type Scope SymbolTable

func Run(sexprTxt string) Evaluable {
	gs := Scope(GlobalSymbolTable())
	sexpr := Parse(sexprTxt)
	if sexpr == nil {
		fmt.Println(PARSE_ERROR)
		return Nil{}
	}
	lispEvaluator := CreateEvaluator(sexpr)
	return Eval(lispEvaluator, gs)
}
