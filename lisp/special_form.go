package lisp

func Define(sym Evaluable, value Evaluable) Evaluable {
	symbol, ok := sym.(Symbol)
	if !ok {
		panic("Type Error")
	}

	symbolTable[symbol.Name] = value
	return symbol
}

// func Lambda(cons Cons) Evaluable {
//  		// attr := cons.Car
//  		f := func(e Cons) Evaluable {
//  			return SExpr{}
//  		}
//  		return Proc(f)
// }
