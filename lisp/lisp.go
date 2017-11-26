package lisp

type Evaluable interface {
	eval(...Scope) Evaluable
}

func Eval(l Evaluable, scs ...Scope) Evaluable {
	return l.eval(scs...)
}

type Scope SymbolTable

func Run(sexprTxt string) Evaluable {
	gs := Scope(GlobalSymbolTable())
	rootAst := Parse(sexprTxt)
	sexpr := CreateEvaluator(rootAst)
	return Eval(sexpr, gs)
}
