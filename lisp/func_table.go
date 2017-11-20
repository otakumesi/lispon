package lisp

type Proc func(Cons) Evaluable

func (f Proc) eval() Evaluable {
	return f
}

func (s Symbol) call(args Cons) Evaluable {
	fn, ok := symbolTable[s].(Proc)
	if !ok {
		panic("Type Error")
	}
	return fn(args)
}

type SymbolTable map[Symbol]Evaluable

var symbolTable = SymbolTable{
	Symbol("+"): Proc(Add),
	Symbol("-"): Proc(Sub),
	Symbol("*"): Proc(Mul),
	Symbol("/"): Proc(Div),
}
