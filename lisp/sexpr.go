package lisp

type SExpr struct {
	symbol Symbol
	lhs    Evaler
	rhs    Evaler
}

type SexprArgs func(*SExpr)

func NewSExpr(sym Symbol, args ...SexprArgs) SExpr {
	sexpr := SExpr{sym, Nil{}, Nil{}}
	for _, arg := range args {
		arg(&sexpr)
	}
	return sexpr
}

func SetLhs(lhs Evaler) SexprArgs {
	return func(s *SExpr) {
		if lhs != nil {
			s.lhs = lhs
		}
	}
}

func SetRhs(rhs Evaler) SexprArgs {
	return func(s *SExpr) {
		if rhs != nil {
			s.rhs = rhs
		}
	}
}

func (s SExpr) eval() Evaler {
	symbol := s.symbol.eval()

	proc, isProc := symbol.(Proc)
	if !isProc {
		return symbol.eval()
	}

	return proc(s.lhs.eval(), s.rhs.eval())
}

func (s SExpr) IsAtom() Evaler {
	return Nil{}
}
