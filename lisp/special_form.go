package lisp

func Define(sym Symbol, value Evaler) Evaler {
	sexpr, isSexpr := value.(SExpr)
	globalSymbolTable := *GlobalSymbolTable()
	if isSexpr {
		globalSymbolTable[sym.Name] = sexpr.eval()
	} else {
		globalSymbolTable[sym.Name] = value
	}
	return sym
}

func Lambda(form SExpr, args ...Symbol) Evaler {
	localSymTable := Scope{}
	f := func(funargs ...Evaler) Evaler {
		GetEnv().Unshift(&localSymTable)
		defer GetEnv().Shift()
		for i, arg := range args {
			localSymTable[arg.Name] = funargs[i]
		}
		return form.eval()
	}
	return Proc(f)
}

func If(condExp, lhsAction, rhsAction Evaler) Evaler {
	if condExp.eval() == Evaler(Nil{}) {
		return rhsAction
	}
	return lhsAction
}
