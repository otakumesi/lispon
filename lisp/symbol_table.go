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

type Bool bool

func (b Bool) eval() Evaluable {
	return b
}

func Eq(lhs, rhs Evaluable) Evaluable {
	return Bool(lhs == rhs)
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

func Define(sym Evaluable, value Evaluable) Evaluable {
	symbol, ok := sym.(Symbol)
	if !ok {
		panic("Type Error")
	}

	symbolTable[symbol.Name] = value
	return Nil{}
}

// func Lambda(cons Cons) Evaluable {
//  		// attr := cons.Car
//  		f := func(e Cons) Evaluable {
//  			return SExpr{}
//  		}
//  		return Proc(f)
// }
