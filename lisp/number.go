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
	case Pair:
		return n.AddPair(rn)
	}
	panic("TypeError")
}

func (n Number) AddNumber(an Number) Evaler {
	return n + an
}

func (n Number) AddPair(c Pair) Evaler {
	re, ok := Car(c).(Number)

	if !ok {
		panic("TypeError")
	}

	sum := n + re

	_, isNil := Cdr(c).(Nil)
	if isNil {
		return sum
	}

	return sum.Add(Cdr(c))
}

func (n Number) Sub(an Evaler) Evaler {
	switch rn := an.(type) {
	case Number:
		return n.SubNumber(rn)
	case Pair:
		return n.SubPair(rn)
	}
	panic("TypeError")
}

func (n Number) SubNumber(an Number) Evaler {
	return n - an
}

func (n Number) SubPair(c Pair) Evaler {
	re, ok := Car(c).(Number)

	if !ok {
		panic("TypeError")
	}

	sum := n - re

	_, isNil := Cdr(c).(Nil)
	if isNil {
		return sum
	}

	return sum.Sub(Cdr(c))
}

func (n Number) Mul(an Evaler) Evaler {
	switch rn := an.(type) {
	case Number:
		return n.MulNumber(rn)
	case Pair:
		return n.MulPair(rn)
	}
	panic("TypeError")
}

func (n Number) MulNumber(on Number) Evaler {
	return n * on
}

func (n Number) MulPair(c Pair) Evaler {
	re, ok := Car(c).(Number)

	if !ok {
		panic("TypeError")
	}

	sum := n * re

	_, isNil := Cdr(c).(Nil)
	if isNil {
		return sum
	}

	return sum.Mul(Cdr(c))
}

func (n Number) Div(an Evaler) Evaler {
	switch rn := an.(type) {
	case Number:
		return n.DivNumber(rn)
	case Pair:
		return n.DivPair(rn)
	}
	panic("TypeError")
}

func (n Number) DivNumber(an Number) Evaler {
	return n / an
}

func (n Number) DivPair(c Pair) Evaler {
	re, ok := Car(c).(Number)

	if !ok {
		panic("TypeError")
	}

	sum := n / re

	_, isNil := Cdr(c).(Nil)
	if isNil {
		return sum
	}

	return sum.Div(Cdr(c))
}
