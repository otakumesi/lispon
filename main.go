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

type OperatorEvaluable interface {
	canOperatorEvaluate() bool
}

type Number float64

func (a Number) eval() Lispable {
	return a
}

func Add(lhs, rhs Lispable) Lispable {
	l, err := lhs.(OperatorEvaluable)
	r, err := rhs.(OperatorEvaluable)
	return lhs + rhs
}

func AddNumber(lhs, rhs Number) Number {
	return lhs + rhs
}

type Symbol string

func (s Symbol) eval() Lispable {
	return s
}

func (s Symbol) toFunc() func(lhs, rhs Lispable) Lispable {
	return Add
}

type Cons struct {
	value Lispable
	cons *Cons
}

func (c Cons) eval() Lispable {
	return c
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
