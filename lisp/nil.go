package lisp

type Nil struct{}

func (n Nil) eval(lss ...LocalScope) Evaluable {
	return n
}
