package lisp

func (s Symbol) call(lhs, rhs Evaluable) Evaluable {
	proc, ok := symbolTable[s.Name].(Proc)
	if !ok {
		panic("Type Error")
	}
	return proc(lhs, rhs)
}

type Scope map[string]Evaluable

var symbolTable = Scope{}

func GlobalSymbolTable() *Scope {
	if len(symbolTable) < 1 {
		symbolTable = Scope{
			"+":  Proc(Add),
			"-":  Proc(Sub),
			"*":  Proc(Mul),
			"/":  Proc(Div),
			"eq": Proc(Eq),
		}
	}
	return &symbolTable
}
