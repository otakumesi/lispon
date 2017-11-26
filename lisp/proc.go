package lisp

type Proc func(Evaluable, Evaluable) Evaluable

func (p Proc) eval(scs ...Scope) Evaluable {
	return p
}
