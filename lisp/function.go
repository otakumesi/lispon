package lisp

import "fmt"

func Cons(args ...Evaler) Evaler {
	if len(args[1:]) < 2 {
		return Pair{args[0], args[1]}
	}
	return Pair{args[0], Cons(args[1])}
}

func Eq(args ...Evaler) Evaler {
	if Cdr(args[1]) == Evaler(Nil{}) {
		return Eq(args[0], Car(args[1]))
	}
	if args[0] != args[1] {
		return Nil{}
	}
	return T{}
}

func Car(args ...Evaler) Evaler {
	if len(args) > 2 {
		return Nil{}
	}
	pair, isPair := args[0].(Pair)

	if !isPair {
		return Nil{}
	}

	return pair.lhs
}

func Cdr(args ...Evaler) Evaler {
	if len(args) > 2 {
		return Nil{}
	}
	pair, isPair := args[0].(Pair)

	if !isPair {
		return Nil{}
	}

	return pair.rhs
}

func Print(args ...Evaler) Evaler {
	fmt.Println(args)
	return Nil{}
}

func IsAtom(args ...Evaler) Evaler {
	if len(args) > 2 || args[1] != Evaler(Nil{}) {
		return Nil{}
	}
	return args[0].IsAtom()
}
