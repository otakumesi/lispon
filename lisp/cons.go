package lisp

type Cons struct {
	Car Evaler
	Cdr Evaler
}

func (c Cons) eval() Evaler {
	return Cons{Car: c.Car.eval(), Cdr: c.Cdr.eval()}
}

func (c Cons) IsAtom() Evaler {
	return Nil{}
}
