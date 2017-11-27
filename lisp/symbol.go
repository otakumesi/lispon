package lisp

type Symbol struct {
	Name     string
	IsQuoted bool
}

type SymbolOption func(*Symbol)

func NewSymbol(name string, opts ...SymbolOption) Symbol {
	sym := Symbol{name, false}
	for _, opt := range opts {
		opt(&sym)
	}
	return sym
}

func SetIsQuote(isQuoted bool) SymbolOption {
	return func(s *Symbol) {
		s.IsQuoted = isQuoted
	}
}

func (s Symbol) eval() Evaluable {
	if s.IsQuoted {
		return String(s.Name)
	}

	for _, sc := range GetEnv().ScopeStacks {
		for name, val := range *sc {
			if name == s.Name {
				return val
			}
		}
	}

	return Nil{}
}
