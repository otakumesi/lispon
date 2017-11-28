package lisp

type T struct{}

func (t T) eval() Evaler {
	return t
}

func (t T) IsAtom() Evaler {
	return t
}

func (t T) String() string {
	return "T"
}

type Nil struct{}

func (n Nil) eval() Evaler {
	return n
}

func (n Nil) IsAtom() Evaler {
	return n
}

func (n Nil) String() string {
	return "Nil"
}
