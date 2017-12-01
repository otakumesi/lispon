package lisp

import "fmt"

type Evaler interface {
	eval() Evaler
	IsAtom() Evaler
}

func Eval(l Evaler) Evaler {
	return l.eval()
}

func Run(sexprTxt string) []Evaler {
	env := GetEnv()
	env.Unshift(GlobalSymbolTable())
	sexprs := Parse(sexprTxt)
	if sexprs == nil {
		fmt.Println(PARSE_ERROR)
		return []Evaler{Nil{}} // TODO あとで例外処理を考えたときに実装する
	}
	lispEvaluators := CreateEvaluators(sexprs)
	var results []Evaler
	for _, evaluator := range lispEvaluators {
		results = append(results, Eval(evaluator))
	}
	return results
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

func (e Env) GetValue(s Symbol) Evaler {
	for _, sc := range GetEnv().Scopes {
		for name, val := range *sc {
			if name == s.Name {
				return val
			}
		}
	}
	return Nil{}
}

var env = &Env{}

func GetEnv() *Env {
	return env
}
