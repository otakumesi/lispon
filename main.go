package main

import "fmt"

type Lispable interface {
	eval() Lispable
}

type SExpr struct {
	symbol Symbol
	lhs    Lispable
	rhs    Lispable
}

func (s SExpr) eval() Lispable {
	symbol := s.symbol
	lhs := s.lhs.eval()
	rhs := s.rhs.eval()
	return symbol.toFunc()(lhs, rhs)
}

type Number float64

func (n Number) eval() Lispable {
	return n
}

type Addable interface {
	Add(Lispable) Lispable
}

func Add(lhs, rhs Lispable) Lispable {
	receiver, ok := lhs.(Addable)
	if !ok {
		panic("TypeError")
	}
	return receiver.Add(rhs)
}

func Sub(lhs, rhs Lispable) Lispable {
	switch l := lhs.(type) {
	case Number:
		return l.Sub(rhs)
	}
	panic("TypeError")
}

func Mul(lhs, rhs Lispable) Lispable {
	switch l := lhs.(type) {
	case Number:
		return l.Mul(rhs)
	}
	panic("TypeError")
}

func Div(lhs, rhs Lispable) Lispable {
	switch l := lhs.(type) {
	case Number:
		return l.Div(rhs)
	}
	panic("TypeError")
}

func (n Number) Add(an Lispable) Lispable {
	switch rn := an.(type) {
	case Number:
		return n + rn
	case Cons:
		num := rn.Car.(Number)
		return n + num.Add(rn.Cdr).(Number)
	case Nil:
		return Number(0)
	}
	panic("TypeError")
}

func (n Number) Sub(an Lispable) Lispable {
	switch rn := an.(type) {
	case Number:
		return n + rn
	case Cons:
		num := rn.Car.(Number)
		return n + num.Sub(rn.Cdr).(Number)
	case Nil:
		return Number(0)
	}
	panic("TypeError")
}

func (n Number) Mul(an Lispable) Lispable {
	switch rn := an.(type) {
	case Number:
		return n * rn
	case Cons:
		num := rn.Car.(Number)
		return n * num.Mul(rn.Cdr).(Number)
	case Nil:
		return Number(1)
	}
	panic("TypeError")
}

func (n Number) Div(an Lispable) Lispable {
	switch rn := an.(type) {
	case Number:
		return n / rn
	case Cons:
		num := rn.Car.(Number)
		return n / num.Div(rn.Cdr).(Number)
	case Nil:
		return Number(1)
	}
	panic("TypeError")
}

type Symbol string

func (s Symbol) eval() Lispable {
	return s
}

func (s Symbol) Add(as Lispable) Lispable {
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

func (s Symbol) toFunc() func(lhs, rhs Lispable) Lispable {
	return funcTable[s]
}

type Cons struct {
	Car Lispable
	Cdr Lispable
}

func (c Cons) Add(ac Lispable) Lispable {
	return Cons{Car: c, Cdr: ac}
}

func (c Cons) eval() Lispable {
	return c
}

func eval(l Lispable) Lispable {
	return l.eval()
}

type Nil struct{}

func (n Nil) eval() Lispable {
	return n
}

type FuncTable map[Symbol]func(Lispable, Lispable) Lispable

var funcTable = FuncTable{
	Symbol("+"): Add,
	Symbol("-"): Sub,
	Symbol("*"): Mul,
	Symbol("/"): Div,
}

func main() {
	sym := Symbol("+")
	lhs := Symbol("otaku")
	cons_l := Symbol("mesi")
	cons_r := Cons{Symbol("!!!"), Nil{}}
	rhs := Cons{cons_l, cons_r}
	lisp := SExpr{sym, lhs, rhs}
	fmt.Println(eval(lisp))
}
