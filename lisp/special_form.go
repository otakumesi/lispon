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
	f := func(lhs, rhs Evaler) Evaler {
		if lhs == Evaler(Nil{}) {
			return form.eval()
		}

		GetEnv().Unshift(&localSymTable)
		localSymTable[args[0].Name] = lhs
		currentRhs := rhs
		for _, arg := range args[1:] {
			argsRhs, isCons := currentRhs.(Cons)
			if isCons {
				localSymTable[arg.Name] = argsRhs.Car
				currentRhs = argsRhs.Cdr
			} else {
				localSymTable[arg.Name] = rhs
				break
			}
		}
		result := form.eval()
		GetEnv().Shift()
		return result
	}
	return Proc(f)
}

func IsAtom(e Evaler) Evaler {
	cons, isCons := e.(Cons)
	if isCons && cons.Cdr == Evaler(Nil{}) {
		return cons.Car.IsAtom()
	}
	return e.IsAtom()
}
