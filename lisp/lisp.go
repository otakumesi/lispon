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
	env.Push(GlobalSymbolTable())
	sexpr := Parse(sexprTxt)
	if sexpr == nil {
		fmt.Println(PARSE_ERROR)
		return Nil{}
	}
	lispEvaluator := CreateEvaluator(sexpr)
	return Eval(lispEvaluator)
}

type Env struct {
	ScopeStacks []*Scope
}

func (e *Env) Push(s *Scope) {
	e.ScopeStacks = append(e.ScopeStacks, s)
}

func (e *Env) Pop() *Scope {
	popScope := e.ScopeStacks[len(e.ScopeStacks)-1]
	e.ScopeStacks = e.ScopeStacks[0 : len(e.ScopeStacks)-2]
	return popScope
}

var env = &Env{}

func GetEnv() *Env {
	return env
}
