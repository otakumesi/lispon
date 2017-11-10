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

type Addable interface {
	Add(Addable) Lispable
}

type Number float64

func (n Number) eval() Lispable {
	return n
}

func (n Number) Add(an Lispable) Lispable {
	switch rn := an.(type) {
	case Number:
		return n + rn
	case Cons:
		num := rn.lhs.(Number)
		return n + num.Add(rn.rhs).(Number)
	case Nil:
		return n
	}
	panic("TypeError")
}

func Add(lhs, rhs Lispable) Lispable {
	switch l := lhs.(type) {
	case Number:
		return l.Add(rhs)
	case Symbol:
		return l.Add(rhs)
	case Cons:
		return l.Add(rhs)
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
		sym := rs.lhs.(Symbol)
		return s + sym.Add(rs.rhs).(Symbol)
	case Nil:
		return s + ""
	}
	panic("TypeError")
}

func (s Symbol) toFunc() func(lhs, rhs Lispable) Lispable {
	return funcTable[s]
}

type Cons struct {
	lhs Lispable
	rhs Lispable
}

func (c Cons) Add(ac Lispable) Lispable {
	return Cons{rhs: c, lhs: ac}
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
}

func main() {
	sym := Symbol("+")
	lhs := Symbol("otaku")
	cons_l := Symbol("mesi")
	cons_r := Symbol("!!!")
	rhs := Cons{cons_l, cons_r}
	lisp := SExpr{sym, lhs, rhs}
	fmt.Println(eval(lisp))
}
