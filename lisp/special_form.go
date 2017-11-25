package lisp

func Define(sym Symbol, value Evaluable) Evaluable {
	sexpr, isSexpr := value.(SExpr)
	if isSexpr {
		symbolTable[sym.Name] = sexpr.eval()
	} else {
		symbolTable[sym.Name] = value
	}
	return sym
}

func Lambda(form SExpr, args ...Symbol) Evaluable {
	localSymTable := LocalScope{}
	f := func(lhs, rhs Evaluable) Evaluable {
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

		return form.eval(localSymTable)
	}
	return Proc(f)
}
