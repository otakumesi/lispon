package lisp

type SExpr struct {
	symbol Symbol
	rhs    Evaluable
	lhs    Evaluable
}

func NewSExpr(sym Symbol, rhs, lhs Evaluable) SExpr {
	return SExpr{sym, rhs, lhs}
}

func (s SExpr) eval() Evaluable {
	proc, ok := s.symbol.eval().(Proc)
	if !ok {
		panic("Type Error")
	}
	return proc(s.rhs, s.lhs)
}
