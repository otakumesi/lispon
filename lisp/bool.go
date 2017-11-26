package lisp

type Bool bool

func (b Bool) eval(scs ...Scope) Evaluable {
	return b
}
