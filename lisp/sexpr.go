package lisp

type SExpr struct {
	symbol Symbol
	lhs    Evaluable
	rhs    Evaluable
}

func NewSExpr(sym Symbol, lhs, rhs Evaluable) SExpr {
	return SExpr{sym, lhs, rhs}
}

func (s SExpr) Eval() Evaluable {
	symbol := s.symbol
	lhs := s.lhs.eval()
	rhs := s.rhs.eval()
	return symbol.toFunc()(lhs, rhs)
}
