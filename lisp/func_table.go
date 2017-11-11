package lisp

type FuncTable map[Symbol]func(Evaluable, Evaluable) Evaluable

var funcTable = FuncTable{
	Symbol("+"): Add,
	Symbol("-"): Sub,
	Symbol("*"): Mul,
	Symbol("/"): Div,
}
