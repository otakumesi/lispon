package lisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLisp(t *testing.T) {
	type Case struct {
		in, out Evaluable
	}
	numAddSexpr := setUpNumberSexpr("+", 5, 6, 7)
	numSubSexpr := setUpNumberSexpr("-", 5, 6, 7)
	numMulSexpr := setUpNumberSexpr("*", 5, 2, 3)
	numDivSexpr := setUpNumberSexpr("/", 15, 5, 3)
	symSexpr := setUpSymbolSexpr("+", "otaku", "mesi", "IO")
	cases := []Case{
		{numAddSexpr.Eval(), Number(18)},
		{numSubSexpr.Eval(), Number(-8)},
		{numMulSexpr.Eval(), Number(30)},
		{numDivSexpr.Eval(), Number(1)},
		{symSexpr.Eval(), Symbol("otakumesiIO")},
	}

	for _, c := range cases {
		assert.Equal(t, c.in, c.out)
	}
}

func setUpNumberSexpr(sym_str string, lhs, rhsR, rhsL int) SExpr {
	sym := Symbol(sym_str)
	car := Number(lhs)
	consL := Number(rhsL)
	consR := Cons{Number(rhsR), Nil{}}
	cdr := Cons{consL, consR}
	sexpr := NewSExpr(sym, car, cdr)
	return sexpr
}

func setUpSymbolSexpr(sym_str, lhs, rhsR, rhsL string) SExpr {
	sym := Symbol(sym_str)
	car := Symbol(lhs)
	consL := Symbol(rhsR)
	consR := Cons{Symbol(rhsL), Nil{}}
	cdr := Cons{consL, consR}
	sexpr := NewSExpr(sym, car, cdr)
	return sexpr
}
