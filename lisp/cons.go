package lisp

type Cons struct {
	Car Evaluable
	Cdr Evaluable
}

func (c Cons) eval(lss ...LocalScope) Evaluable {
	return Cons{Car: c.Car.eval(lss...), Cdr: c.Cdr.eval(lss...)}
}
