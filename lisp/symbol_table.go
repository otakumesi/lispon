package lisp

func (s Symbol) call(lhs, rhs Evaluable) Evaluable {
	proc, ok := symbolTable[s.Name].(Proc)
	if !ok {
		panic("Type Error")
	}
	return proc(lhs, rhs)
}

type SymbolTable map[string]Evaluable

var symbolTable = SymbolTable{}

func GlobalSymbolTable() SymbolTable {
	if len(symbolTable) < 1 {
		symbolTable = SymbolTable{
			"+":  Proc(Add),
			"-":  Proc(Sub),
			"*":  Proc(Mul),
			"/":  Proc(Div),
			"eq": Proc(Eq),
		}
	}
	return symbolTable
}
