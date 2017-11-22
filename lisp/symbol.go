package lisp

type Symbol struct {
	Name     String
	IsQuoted bool
}

func (s Symbol) eval() Evaluable {
	if s.IsQuoted {
		return s.Name
	}
	return symbolTable[s.Name]
}
