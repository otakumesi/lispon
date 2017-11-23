package lisp

type Evaluable interface {
	eval() Evaluable
}

func Eval(l Evaluable) Evaluable {
	return l.eval()
}
