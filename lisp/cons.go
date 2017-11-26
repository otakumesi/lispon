package lisp

type Cons struct {
	Car Evaluable
	Cdr Evaluable
}

func (c Cons) eval(scs ...Scope) Evaluable {
	return Cons{Car: c.Car.eval(scs...), Cdr: c.Cdr.eval(scs...)}
}
