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

func (a Number) eval() Lispable {
	return a
}

type Symbol string

func (s Symbol) eval() Lispable {
	return s
}

func (s Symbol) toFunc() func(lhs, rhs Lispable) Lispable {
	return Add
}

func Add(lhs, rhs Lispable) Lispable {
	l := lhs.(Number)
	r := rhs.(Number)
	return AddNumber(l, r)
}

func AddNumber(lhs, rhs Number) Number {
	return lhs + rhs
}

func eval(l Lispable) Lispable {
	return l.eval()
}

func main() {
	sym := Symbol("plus")
	lhs := Number(3)
	rhs := Number(5)
	lisp := SExpr{sym, lhs, rhs}
	fmt.Println(eval(lisp))
}
