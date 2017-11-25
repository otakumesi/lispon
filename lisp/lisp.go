package lisp

type Evaluable interface {
	eval(...LocalScope) Evaluable
}

func Eval(l Evaluable) Evaluable {
	return l.eval()
}

type LocalScope SymbolTable

func Run(sexprTxt string) Evaluable {
	rootAst := Parse(sexprTxt)
	sexpr := CreateEvaluator(rootAst)
	return Eval(sexpr)
}
