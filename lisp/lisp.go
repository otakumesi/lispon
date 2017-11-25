package lisp

type Evaluable interface {
	eval(...LocalScope) Evaluable
}

func Eval(l Evaluable) Evaluable {
	return l.eval()
}

type LocalScope SymbolTable
