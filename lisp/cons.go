package lisp

type Cons struct {
	Car Evaluable
	Cdr Evaluable
}

func (c Cons) Add(ac Evaluable) Evaluable {
	return Cons{Car: c, Cdr: ac}
}

func (c Cons) eval() Evaluable {
	return c
}
