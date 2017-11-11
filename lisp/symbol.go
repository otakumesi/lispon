package lisp

type Symbol string

func (s Symbol) eval() Evaluable {
	return s
}

func (s Symbol) Add(as Evaluable) Evaluable {
	switch rs := as.(type) {
	case Symbol:
		return s + rs
	case Cons:
		sym := rs.Car.(Symbol)
		return s + sym.Add(rs.Cdr).(Symbol)
	case Nil:
		return s + ""
	}
	panic("TypeError")
}

func (s Symbol) toFunc() func(lhs, rhs Evaluable) Evaluable {
	return funcTable[s]
}
