package lisp

type Bool bool

func (b Bool) eval(lss ...LocalScope) Evaluable {
	return b
}
