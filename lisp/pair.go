package lisp

type Pair struct {
	lhs Evaler
	rhs Evaler
}

func (p Pair) eval() Evaler {
	return Pair{lhs: Car(p).eval(), rhs: Cdr(p).eval()}
}

func (p Pair) IsAtom() Evaler {
	return Nil{}
}
