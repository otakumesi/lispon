package lisp

import "fmt"

type Evaluable interface {
	eval() Evaluable
}

func Eval(l Evaluable) Evaluable {
	return l.eval()
}

func Run(sexprTxt string) Evaluable {
	env := GetEnv()
	env.Unshift(GlobalSymbolTable())
	sexpr := Parse(sexprTxt)
	if sexpr == nil {
		fmt.Println(PARSE_ERROR)
		return Nil{}
	}
	lispEvaluator := CreateEvaluator(sexpr)
	return Eval(lispEvaluator)
}

type Env struct {
	Scopes []*Scope
}

func (e *Env) Unshift(s *Scope) {
	e.Scopes = append([]*Scope{s}, e.Scopes...)
}

func (e *Env) Shift() *Scope {
	shiftScope, newScopes := e.Scopes[0], e.Scopes[1:]
	e.Scopes = newScopes
	return shiftScope
}

var env = &Env{}

func GetEnv() *Env {
	return env
}
