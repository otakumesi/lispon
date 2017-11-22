package lisp

type SExpr struct {
	symbol Symbol
	args   Cons
}

func NewSExpr(sym Symbol, args Cons) SExpr {
	return SExpr{sym, args}
}

func (s SExpr) eval() Evaluable {
	proc, ok := s.symbol.eval().(Proc)
	if !ok {
		panic("Type Error")
	}
	return proc(s.args)
}
