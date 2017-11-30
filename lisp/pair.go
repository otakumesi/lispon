package lisp

type Pair struct {
	lhs Evaler
	rhs Evaler
}

func (p Pair) eval() Evaler {
	return Pair{lhs: p.Car().eval(), rhs: p.Cdr().eval()}
}

func (p Pair) Car() Evaler {
	return p.lhs
}

func (p Pair) Cdr() Evaler {
	return p.rhs
}

func (p Pair) IsAtom() Evaler {
	return Nil{}
}
