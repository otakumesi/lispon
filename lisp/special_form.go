package lisp

func Define(sym Symbol, value Evaluable) Evaluable {
	sexpr, isSexpr := value.(SExpr)
	globalSymbolTable := *GlobalSymbolTable()
	if isSexpr {
		globalSymbolTable[sym.Name] = sexpr.eval()
	} else {
		globalSymbolTable[sym.Name] = value
	}
	return sym
}

func Lambda(form SExpr, args ...Symbol) Evaluable {
	localSymTable := Scope{}
	f := func(lhs, rhs Evaluable) Evaluable {
		if lhs == Evaluable(Nil{}) {
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
