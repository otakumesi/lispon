package lisp

type SExpr struct {
	symbol Symbol
	lhs    Evaluable
	rhs    Evaluable
}

type SexprArgs func(*SExpr)

func NewSExpr(sym Symbol, args ...SexprArgs) SExpr {
	sexpr := SExpr{sym, Nil{}, Nil{}}
	for _, arg := range args {
		arg(&sexpr)
	}
	return sexpr
}

func SetLhs(lhs Evaluable) SexprArgs {
	return func(s *SExpr) {
		if lhs != nil {
			s.lhs = lhs
		}
	}
}

func SetRhs(rhs Evaluable) SexprArgs {
	return func(s *SExpr) {
		if rhs != nil {
			s.rhs = rhs
		}
	}
}

func (s SExpr) eval(scs ...Scope) Evaluable {
	symbol := s.symbol.eval(scs...)

	proc, isProc := symbol.(Proc)
	if !isProc {
		return symbol.eval(scs...)
	}

	return proc(s.lhs.eval(scs...), s.rhs.eval(scs...))
}
