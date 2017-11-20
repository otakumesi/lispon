package lisp

type SExpr struct {
	symbol Symbol
	args   Cons
}

func NewSExpr(sym Symbol, args Cons) SExpr {
	return SExpr{sym, args}
}

func (s SExpr) eval() Evaluable {
	symbol := s.symbol
	return symbol.call(s.args)
}
