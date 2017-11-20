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
		return s.AddCons(rs)
	case Nil:
		return s + ""
	}
	panic("TypeError")
}

func (s Symbol) AddSym(as Symbol) Evaluable {
	return s + as
}

func (s Symbol) AddCons(c Cons) Evaluable {
	Car, ok := c.Car.(Symbol)

	if !ok {
		panic("TypeError")
	}

	result := s + Car

	_, isNil := c.Cdr.(Nil)
	if isNil {
		return result
	}

	return result.Add(c.Cdr)
}
