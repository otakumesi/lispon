package lisp

type Proc func(Cons) Evaluable

func (f Proc) eval() Evaluable {
	return f
}

func (s Symbol) call(args Cons) Evaluable {
	fn, ok := symbolTable[s.Name].(Proc)
	if !ok {
		panic("Type Error")
	}
	return fn(args)
}

type Bool bool

func (b Bool) eval() Evaluable {
	return b
}

func Eq(cons Cons) Evaluable {
	return Bool(cons.Car == cons.Cdr)
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
		String("lambda"): Proc(Lambda),
	}
}

func Define(cons Cons) Evaluable {
	symbol, ok := cons.Car.(Symbol)
	if !ok {
		panic("Type Error")
	}

	symbolTable[symbol.Name] = cons
	return Nil{}
}

func Lambda(cons Cons) Evaluable {
	// attr := cons.Car
	f := func(e Cons) Evaluable {
		return SExpr{}
	}
	return Proc(f)
}
