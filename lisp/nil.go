package lisp

type Nil struct{}

func (n Nil) eval(scs ...Scope) Evaluable {
	return n
}
