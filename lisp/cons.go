package lisp

type Cons struct {
	car Evaluable
	cdr Evaluable
}

func (c Cons) Add(ac Evaluable) Evaluable {
	return Cons{car: c, cdr: ac}
}

func (c Cons) eval() Evaluable {
	return c
}
