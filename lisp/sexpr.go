package lisp

type SExpr struct {
	symbol Symbol
	lhs    Evaluable
	rhs    Evaluable
}

func NewSExpr(sym Symbol, args ...Args) SExpr {
	sexpr := SExpr{sym, Nil{}, Nil{}}
	for _, arg := range args {
		arg(&sexpr)
	}
	return sexpr
}

type Args func(*SExpr)

func SetLhs(lhs Evaluable) Args {
	return func(s *SExpr) {
		if lhs != nil {
			s.lhs = lhs
		}
	}
}

func SetRhs(rhs Evaluable) Args {
	return func(s *SExpr) {
		if rhs != nil {
			s.rhs = rhs
		}
	}
}

func (s SExpr) eval(lss ...LocalScope) Evaluable {
	symbol := s.symbol.eval(lss...)
	proc, isProc := symbol.(Proc)
	if !isProc {
		return symbol.eval(lss...)
	}

	return proc(s.lhs.eval(lss...), s.rhs.eval(lss...))
}
