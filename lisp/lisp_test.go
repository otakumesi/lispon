package lisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLisp(t *testing.T) {
	type Case struct {
		in, out Evaler
	}
	numAddSexpr := setUpNumberSexpr("+", 5, 6, 7)
	numSubSexpr := setUpNumberSexpr("-", 5, 6, 7)
	numMulSexpr := setUpNumberSexpr("*", 5, 2, 3)
	numDivSexpr := setUpNumberSexpr("/", 15, 5, 3)
	strSexpr := setUpStringSexpr("+", "otaku", "mesi", "IO")
	env := GetEnv()
	env.Unshift(GlobalSymbolTable())

	cases := []Case{
		{Eval(numAddSexpr), Number(18)},
		{Eval(numSubSexpr), Number(-8)},
		{Eval(numMulSexpr), Number(30)},
		{Eval(numDivSexpr), Number(1)},
		{Eval(strSexpr), String("otakumesiIO")},
	}

	for _, c := range cases {
		assert.Equal(t, c.in, c.out)
	}
}

func setUpNumberSexpr(sym_str string, args ...int) SExpr {
	sym := NewSymbol(sym_str)
	pair := Pair{Number(args[1]), Number(args[2])}
	sexpr := NewSExpr(sym, SetLhs(Number(args[0])), SetRhs(pair))
	return sexpr
}

func setUpStringSexpr(sym_str string, args ...string) SExpr {
	sym := NewSymbol(sym_str)
	pair := Pair{String(args[1]), String(args[2])}
	sexpr := NewSExpr(sym, SetLhs(String(args[0])), SetRhs(pair))
	return sexpr
}
