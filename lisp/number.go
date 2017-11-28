package lisp

type Number float64

func (n Number) eval() Evaler {
	return n
}

func (n Number) IsAtom() Evaler {
	return T{}
}

func (n Number) Add(an Evaler) Evaler {
	switch rn := an.(type) {
	case Number:
		return n.AddNumber(rn)
	case Cons:
		return n.AddCons(rn)
	}
	panic("TypeError")
}

func (n Number) AddNumber(an Number) Evaler {
	return n + an
}

func (n Number) AddCons(c Cons) Evaler {
	re, ok := c.Car.(Number)

	if !ok {
		panic("TypeError")
	}

	sum := n + re

	_, isNil := c.Cdr.(Nil)
	if isNil {
		return sum
	}

	return sum.Add(c.Cdr)
}

func (n Number) Sub(an Evaler) Evaler {
	switch rn := an.(type) {
	case Number:
		return n.SubNumber(rn)
	case Cons:
		return n.SubCons(rn)
	}
	panic("TypeError")
}

func (n Number) SubNumber(an Number) Evaler {
	return n - an
}

func (n Number) SubCons(c Cons) Evaler {
	re, ok := c.Car.(Number)

	if !ok {
		panic("TypeError")
	}

	sum := n - re

	_, isNil := c.Cdr.(Nil)
	if isNil {
		return sum
	}

	return sum.Sub(c.Cdr)
}

func (n Number) Mul(an Evaler) Evaler {
	switch rn := an.(type) {
	case Number:
		return n.MulNumber(rn)
	case Cons:
		return n.MulCons(rn)
	}
	panic("TypeError")
}

func (n Number) MulNumber(on Number) Evaler {
	return n * on
}

func (n Number) MulCons(c Cons) Evaler {
	re, ok := c.Car.(Number)

	if !ok {
		panic("TypeError")
	}

	sum := n * re

	_, isNil := c.Cdr.(Nil)
	if isNil {
		return sum
	}

	return sum.Mul(c.Cdr)
}

func (n Number) Div(an Evaler) Evaler {
	switch rn := an.(type) {
	case Number:
		return n.DivNumber(rn)
	case Cons:
		return n.DivCons(rn)
	}
	panic("TypeError")
}

func (n Number) DivNumber(an Number) Evaler {
	return n / an
}

func (n Number) DivCons(c Cons) Evaler {
	re, ok := c.Car.(Number)

	if !ok {
		panic("TypeError")
	}

	sum := n / re

	_, isNil := c.Cdr.(Nil)
	if isNil {
		return sum
	}

	return sum.Div(c.Cdr)
}
