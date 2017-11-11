package lisp

type Symbol string

func (s Symbol) eval() Evaluable {
	return s
}

func (s Symbol) Add(as Evaluable) Evaluable {
	switch rs := as.(type) {
	case Symbol:
		return s.AddSym(rs)
	case Cons:
	case Nil:
		return s + ""
	}
	panic("TypeError")
}

func (s Symbol) AddSym(as Symbol) Evaluable {
	return s + as
}

func (s Symbol) AddCons(c Cons) Evaluable {
	car, ok := c.car.(Symbol)

	if !ok {
		panic("TypeError")
	}

	result := s + car

	_, isNil := c.cdr.(Nil)
	if isNil {
		return result
	}

	return result.Add(c.cdr)
}

func (s Symbol) toFunc() func(lhs, rhs Evaluable) Evaluable {
	return funcTable[s]
}
