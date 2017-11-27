package lisp

type Nil struct{}

func (n Nil) eval() Evaluable {
	return n
}
