package lisp

type Proc func(Evaluable, Evaluable) Evaluable

func (p Proc) eval(lss ...LocalScope) Evaluable {
	return p
}
