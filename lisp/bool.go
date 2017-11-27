package lisp

type Bool bool

func (b Bool) eval() Evaluable {
	return b
}
