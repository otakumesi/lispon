package lisp

func Define(sym Symbol, value Evaler) Evaler {
	sexpr, isSexpr := value.(SExpr)

	symbolTable := *GetEnv().Scopes[0]
	if isSexpr {
		symbolTable[sym.Name] = sexpr.eval()
	} else {
		symbolTable[sym.Name] = value
	}
	return sym
}

func Lambda(forms []SExpr, args ...Symbol) Evaler {
	localSymTable := Scope{}
	f := func(funargs ...Evaler) Evaler {
		GetEnv().Unshift(&localSymTable)
		defer GetEnv().Shift()
		for i, arg := range args {
			localSymTable[arg.Name] = funargs[i]
		}
		var results []Evaler
		for _, form := range forms {
			results = append(results, form.eval())
		}
		return results[len(results)-1]
	}
	return Proc(f)
}

func If(condExp, lhsAction, rhsAction Evaler) Evaler {
	if condExp.eval() == Evaler(Nil{}) {
		return rhsAction
	}
	return lhsAction
}

func Quote(args ...Evaler) Evaler {
	if len(args) > 1 {
		return Cons(args...)
	}
	sym, isSym := args[0].(Symbol)
	if isSym {
		return NewSymbol(sym.Name, SetIsQuote(true))
	}
	return Nil{}
}
