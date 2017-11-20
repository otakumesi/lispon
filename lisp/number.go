package lisp

type Number float64

func (n Number) eval() Evaluable {
	return n
}

func (n Number) Add(e Evaluable) Evaluable {
	switch re := e.(type) {
	case Number:
		return n.AddNumber(re)
	case Cons:
		return n.AddCons(re)
	}
	panic("TypeError")
}

func (n Number) AddNumber(on Number) Evaluable {
	return n + on
}

func (n Number) AddCons(c Cons) Evaluable {
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

func (n Number) Sub(an Evaluable) Evaluable {
	switch rn := an.(type) {
	case Number:
		return n.SubNumber(rn)
	case Cons:
		return n.SubCons(rn)
	}
	panic("TypeError")
}

func (n Number) SubNumber(on Number) Evaluable {
	return n - on
}

func (n Number) SubCons(c Cons) Evaluable {
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

func (n Number) Mul(an Evaluable) Evaluable {
	switch rn := an.(type) {
	case Number:
		return n.MulNumber(rn)
	case Cons:
		return n.MulCons(rn)
	}
	panic("TypeError")
}

func (n Number) MulNumber(on Number) Evaluable {
	return n * on
}

func (n Number) MulCons(c Cons) Evaluable {
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

func (n Number) Div(an Evaluable) Evaluable {
	switch rn := an.(type) {
	case Number:
		return n.DivNumber(rn)
	case Cons:
		return n.DivCons(rn)
	}
	panic("TypeError")
}

func (n Number) DivNumber(on Number) Evaluable {
	return n * on
}

func (n Number) DivCons(c Cons) Evaluable {
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
