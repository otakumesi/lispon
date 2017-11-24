package lisp

type Adder interface {
	Add(Evaluable) Evaluable
}

func Add(lhs, rhs Evaluable) Evaluable {
	sym, isSym := lhs.(Symbol)
	var adder Evaluable
	if isSym {
		adder = sym.eval()
	} else {
		adder = lhs
	}
	receiver, ok := adder.(Adder)
	if !ok {
		panic("TypeError")
	}
	return receiver.Add(rhs)
}

type Suber interface {
	Sub(Evaluable) Evaluable
}

func Sub(lhs, rhs Evaluable) Evaluable {
	sym, isSym := lhs.(Symbol)
	var suber Evaluable
	if isSym {
		suber = sym.eval()
	} else {
		suber = lhs
	}
	receiver, ok := suber.(Suber)
	if !ok {
		panic("TypeError")
	}
	return receiver.Sub(rhs)
}

type Muler interface {
	Mul(Evaluable) Evaluable
}

func Mul(lhs, rhs Evaluable) Evaluable {
	sym, isSym := lhs.(Symbol)
	var muler Evaluable
	if isSym {
		muler = sym.eval()
	} else {
		muler = lhs
	}
	receiver, ok := muler.(Muler)
	if !ok {
		panic("TypeError")
	}
	return receiver.Mul(rhs)
}

type Diver interface {
	Div(Evaluable) Evaluable
}

func Div(lhs, rhs Evaluable) Evaluable {
	sym, isSym := lhs.(Symbol)
	var diver Evaluable
	if isSym {
		diver = sym.eval()
	} else {
		diver = lhs
	}
	receiver, ok := diver.(Diver)
	if !ok {
		panic("TypeError")
	}
	return receiver.Div(rhs)
}

func Eq(lhs, rhs Evaluable) Evaluable {
	cons, ok := rhs.(Cons)
	if ok && (cons.Cdr == Nil{}) {
		return Eq(lhs, cons.Car)
	}
	return Bool(lhs == rhs)
}
