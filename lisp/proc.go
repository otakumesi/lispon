package lisp

type Proc func(Evaluable, Evaluable) Evaluable

func (p Proc) eval() Evaluable {
	return p
}
