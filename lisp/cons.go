package lisp

type Cons struct {
	Car Evaluable
	Cdr Evaluable
}

func (c Cons) eval() Evaluable {
	return Cons{Car: c.Car.eval(), Cdr: c.Cdr.eval()}
}
