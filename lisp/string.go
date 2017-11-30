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
	case Pair:
		return s.AddPair(rs)
	case Nil:
		return s + ""
	}
	panic("TypeError")
}

func (s String) AddStr(as String) Evaler {
	return s + as
}

func (s String) AddPair(c Pair) Evaler {
	car, ok := Car(c).(String)

	if !ok {
		panic("TypeError")
	}

	result := s + car

	_, isNil := Cdr(c).(Nil)
	if isNil {
		return result
	}

	return result.Add(Cdr(c))
}
