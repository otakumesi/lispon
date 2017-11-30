package lisp

type Adder interface {
	Add(Evaler) Evaler
}

func Add(args ...Evaler) Evaler {
	receiver, ok := args[0].(Adder)
	if !ok {
		panic("TypeError")
	}
	return receiver.Add(args[1])
}

type Suber interface {
	Sub(Evaler) Evaler
}

func Sub(args ...Evaler) Evaler {
	receiver, ok := args[0].(Suber)
	if !ok {
		panic("TypeError")
	}
	return receiver.Sub(args[1])
}

type Muler interface {
	Mul(Evaler) Evaler
}

func Mul(args ...Evaler) Evaler {
	receiver, ok := args[0].(Muler)
	if !ok {
		panic("TypeError")
	}
	return receiver.Mul(args[1])
}

type Diver interface {
	Div(Evaler) Evaler
}

func Div(args ...Evaler) Evaler {
	receiver, ok := args[0].(Diver)
	if !ok {
		panic("TypeError")
	}
	return receiver.Div(args[1])
}
