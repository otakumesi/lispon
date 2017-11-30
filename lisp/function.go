package lisp

import "fmt"

func Cons(lhs, rhs Evaler) Evaler {
	return Pair{lhs, rhs}
}

func Eq(lhs, rhs Evaler) Evaler {
	cons, ok := rhs.(Pair)
	if ok && (cons.Cdr() == Nil{}) {
		return Eq(lhs, cons.Car())
	}
	if lhs != rhs {
		return Nil{}
	}
	return T{}
}

func Print(lhs, rhs Evaler) Evaler {
	if rhs != Evaler(Nil{}) {
		fmt.Println(lhs, rhs)
	} else {
		fmt.Println(lhs)
	}
	return Nil{}
}
