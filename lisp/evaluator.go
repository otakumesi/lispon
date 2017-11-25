package lisp

import (
	"strconv"

	parsec "github.com/prataprc/goparsec"
)

func CreateEvaluator(ast parsec.Queryable) Evaluable {
	children := ast.GetChildren()
	switch ast.GetName() {
	case "sexpr":
		return CreateEvaluator(ast.GetChildren()[0])
	case "defineExpr":
		return createDefine(children)
	case "lambdaExpr":
		return createLambda(children)
	case "expr":
		return createExpr(children)
	case "items":
		return createItems(children)
	case "item":
		return createItem(children)
	case "INT":
		i, err := strconv.Atoi(ast.GetValue())
		if err != nil {
			panic(err)
		}
		return Number(i)
	case "FLOAT":
		f, err := strconv.ParseFloat(ast.GetValue(), 64)
		if err != nil {
			panic(err)
		}
		return Number(f)
	case "STRING":
		rawStr := ast.GetValue()
		return String(rawStr[1 : len(rawStr)-1])
	case "IDENT":
		return NewSymbol(ast.GetValue())
	case "QUOTED_SYMBOL":
		rawStr := ast.GetValue()
		return NewSymbol(rawStr[1:len(rawStr)], SetIsQuote(true))
	}
	return Nil{}
}

func createItems(children []parsec.Queryable) Evaluable {
	if children[1].GetName() == "missing" {
		return CreateEvaluator(children[0])
	}
	return Cons{CreateEvaluator(children[0]), CreateEvaluator(children[1])}
}

func createItem(children []parsec.Queryable) Evaluable {
	return CreateEvaluator(children[0])
}

func createExpr(children []parsec.Queryable) SExpr {
	sym := NewSymbol(children[1].GetValue())

	if len(children) < 2 {
		return NewSExpr(sym)
	}
	lhs := CreateEvaluator(children[2])

	if len(children) < 3 {
		return NewSExpr(sym, SetLhs(lhs))
	}
	rhs := CreateEvaluator(children[3])

	return NewSExpr(sym, SetLhs(lhs), SetRhs(rhs))
}

func createDefine(children []parsec.Queryable) Evaluable {
	argSym, ok := CreateEvaluator(children[2]).(Symbol)

	if !ok {
		panic(RUNTIME_ERROR)
	}

	expr := CreateEvaluator(children[3])
	return Define(argSym, expr)
}

func createLambda(children []parsec.Queryable) Evaluable {
	form, ok := CreateEvaluator(children[5]).(SExpr)

	if !ok {
		panic(RUNTIME_ERROR)
	}

	var args []Symbol
	for _, child := range children[3].GetChildren() {
		arg := CreateEvaluator(child)
		if arg == Evaluable(Nil{}) {
			continue
		}

		argSym, ok := arg.(Symbol)
		if !ok {
			panic(RUNTIME_ERROR)
		}

		args = append(args, argSym)
	}
	return Lambda(form, args...)
}
