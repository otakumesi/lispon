package lisp

type String string

func (s String) eval(scs ...Scope) Evaluable {
	return s
}

func (s String) Add(as Evaluable) Evaluable {
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

func (s String) AddStr(as String) Evaluable {
	return s + as
}

func (s String) AddCons(c Cons) Evaluable {
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
