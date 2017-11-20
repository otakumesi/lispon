package lisp

type Evaluable interface {
	eval() Evaluable
}

type Adder interface {
	Add(Evaluable) Evaluable
}

func Add(cons Cons) Evaluable {
	receiver, ok := cons.Car.(Adder)
	if !ok {
		panic("TypeError")
	}
	return receiver.Add(cons.Cdr)
}

type Suber interface {
	Sub(Evaluable) Evaluable
}

func Sub(cons Cons) Evaluable {
	receiver, ok := cons.Car.(Suber)
	if !ok {
		panic("TypeError")
	}
	return receiver.Sub(cons.Cdr)
}

type Muler interface {
	Mul(Evaluable) Evaluable
}

func Mul(cons Cons) Evaluable {
	receiver, ok := cons.Car.(Muler)
	if !ok {
		panic("TypeError")
	}
	return receiver.Mul(cons.Cdr)
}

type Diver interface {
	Div(Evaluable) Evaluable
}

func Div(cons Cons) Evaluable {
	receiver, ok := cons.Car.(Diver)
	if !ok {
		panic("TypeError")
	}
	return receiver.Div(cons.Cdr)
}

func Eval(l Evaluable) Evaluable {
	return l.eval()
}

type Nil struct{}

func (n Nil) eval() Evaluable {
	return n
}
