package lisp

type Proc func(Evaluable, Evaluable) Evaluable

func (f Proc) eval() Evaluable {
	return f
}

func (s Symbol) call(lhs, rhs Evaluable) Evaluable {
	fn, ok := symbolTable[s.Name].(Proc)
	if !ok {
		panic("Type Error")
	}
	return fn(lhs, rhs)
}

type SymbolTable map[String]Evaluable

var symbolTable = SymbolTable{}

func Init() {
	symbolTable = SymbolTable{
		String("+"):      Proc(Add),
		String("-"):      Proc(Sub),
		String("*"):      Proc(Mul),
		String("/"):      Proc(Div),
		String("eq"):     Proc(Eq),
		String("define"): Proc(Define),
		// String("lambda"): Proc(Lambda),
	}
}
