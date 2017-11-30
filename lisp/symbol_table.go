package lisp

type Scope map[string]Evaler

var symbolTable = Scope{}

func GlobalSymbolTable() *Scope {
	if len(symbolTable) < 1 {
		symbolTable = Scope{
			"+":     Proc(Add),
			"-":     Proc(Sub),
			"*":     Proc(Mul),
			"/":     Proc(Div),
			"eq":    Proc(Eq),
			"cons":  Proc(Cons),
			"car":   Proc(Car),
			"cdr":   Proc(Cdr),
			"print": Proc(Print),
			"atom":  Proc(IsAtom),
		}
	}
	return &symbolTable
}
