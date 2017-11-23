package lisp

type Evaluable interface {
	eval() Evaluable
}

type Adder interface {
	Add(Evaluable) Evaluable
}

func Add(lhs, rhs Evaluable) Evaluable {
	receiver, ok := lhs.(Adder)
	if !ok {
		panic("TypeError")
	}
	return receiver.Add(rhs)
}

type Suber interface {
	Sub(Evaluable) Evaluable
}

func Sub(lhs, rhs Evaluable) Evaluable {
	receiver, ok := lhs.(Suber)
	if !ok {
		panic("TypeError")
	}
	return receiver.Sub(rhs)
}

type Muler interface {
	Mul(Evaluable) Evaluable
}

func Mul(lhs, rhs Evaluable) Evaluable {
	receiver, ok := lhs.(Muler)
	if !ok {
		panic("TypeError")
	}
	return receiver.Mul(rhs)
}

type Diver interface {
	Div(Evaluable) Evaluable
}

func Div(lhs, rhs Evaluable) Evaluable {
	receiver, ok := lhs.(Diver)
	if !ok {
		panic("TypeError")
	}
	return receiver.Div(rhs)
}

func Eval(l Evaluable) Evaluable {
	return l.eval()
}

type Nil struct{}

func (n Nil) eval() Evaluable {
	return n
}
