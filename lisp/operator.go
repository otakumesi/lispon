package lisp

import "fmt"

type Adder interface {
	Add(Evaler) Evaler
}

func Add(lhs, rhs Evaler) Evaler {
	receiver, ok := lhs.(Adder)
	if !ok {
		panic("TypeError")
	}
	return receiver.Add(rhs)
}

type Suber interface {
	Sub(Evaler) Evaler
}

func Sub(lhs, rhs Evaler) Evaler {
	receiver, ok := lhs.(Suber)
	if !ok {
		panic("TypeError")
	}
	return receiver.Sub(rhs)
}

type Muler interface {
	Mul(Evaler) Evaler
}

func Mul(lhs, rhs Evaler) Evaler {
	receiver, ok := lhs.(Muler)
	if !ok {
		panic("TypeError")
	}
	return receiver.Mul(rhs)
}

type Diver interface {
	Div(Evaler) Evaler
}

func Div(lhs, rhs Evaler) Evaler {
	receiver, ok := lhs.(Diver)
	if !ok {
		panic("TypeError")
	}
	return receiver.Div(rhs)
}

func Eq(lhs, rhs Evaler) Evaler {
	cons, ok := rhs.(Cons)
	if ok && (cons.Cdr == Nil{}) {
		return Eq(lhs, cons.Car)
	}
	if lhs != rhs {
		return Nil{}
	}
	fmt.Println(T{})
	return T{}
}
