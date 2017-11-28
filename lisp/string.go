package lisp

type String string

func (s String) eval() Evaler {
	return s
}

func (s String) IsAtom() Evaler {
	return T{}
}

func (s String) Add(as Evaler) Evaler {
	switch rs := as.(type) {
	case String:
		return s.AddStr(rs)
	case Cons:
		return s.AddCons(rs)
	case Nil:
		return s + ""
	}
	panic("TypeError")
}

func (s String) AddStr(as String) Evaler {
	return s + as
}

func (s String) AddCons(c Cons) Evaler {
	Car, ok := c.Car.(String)

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
