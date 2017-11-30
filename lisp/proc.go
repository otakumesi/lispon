package lisp

type Proc func(...Evaler) Evaler

func (p Proc) eval() Evaler {
	return p
}

func (p Proc) IsAtom() Evaler {
	return T{}
}
