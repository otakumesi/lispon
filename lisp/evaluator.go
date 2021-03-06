package lisp

import (
	"strconv"

	parsec "github.com/prataprc/goparsec"
)

// CreateEvaluators parse results to Evaluators
func CreateEvaluators(ast parsec.Queryable) []Evaler {
	var evaluators []Evaler
	if ast.GetName() == "sexprs" {
		for _, sexpr := range ast.GetChildren() {
			evaluators = append(evaluators, CreateEvaluator(sexpr))
		}
	}
	return evaluators
}

// CreateEvaluator parse result to Evaluator
func CreateEvaluator(ast parsec.Queryable) Evaler {
	children := ast.GetChildren()
	switch ast.GetName() {
	case "sexpr":
		return CreateEvaluator(ast.GetChildren()[0])
	case "defineExpr":
		return createDefine(children)
	case "lambdaExpr":
		return createLambda(children)
	case "pipeExpr":
		return createPipe(children)
	case "ifExpr":
		return createIf(children)
	case "quoteExpr":
		return createQuote(children)
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
	case "OPERATOR":
		return NewSymbol(ast.GetValue())
	case "QUOTED_SYMBOL":
		rawStr := ast.GetValue()
		return NewSymbol(rawStr[1:len(rawStr)], SetIsQuote(true))
	}
	return Nil{}
}

func createItems(children []parsec.Queryable) Evaler {
	if children[1].GetName() == "missing" {
		return CreateEvaluator(children[0])
	}
	return Pair{CreateEvaluator(children[0]), CreateEvaluator(children[1])}
}

func createItem(children []parsec.Queryable) Evaler {
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

func createDefine(children []parsec.Queryable) Evaler {
	argSym, ok := CreateEvaluator(children[2]).(Symbol)

	if !ok {
		panic(RUNTIME_ERROR)
	}

	expr := CreateEvaluator(children[3])
	return Define(argSym, expr)
}

func createLambda(children []parsec.Queryable) Evaler {
	var forms []SExpr
	for _, child := range children[5].GetChildren() {
		form, ok := CreateEvaluator(child).(SExpr)

		if !ok {
			panic(RUNTIME_ERROR)
		}

		forms = append(forms, form)
	}

	var args []Symbol
	for _, child := range children[3].GetChildren() {
		arg := CreateEvaluator(child)
		if arg == Evaler(Nil{}) {
			continue
		}

		argSym, ok := arg.(Symbol)
		if !ok {
			panic(RUNTIME_ERROR)
		}

		args = append(args, argSym)
	}
	return Lambda(forms, args...)
}

func createPipe(children []parsec.Queryable) Evaler {
	lhs := CreateEvaluator(children[1])

	pipeTarget := children[3].GetChildren()
	sym, ok := CreateEvaluator(pipeTarget[0]).(Symbol)
	if !ok {
		panic(RUNTIME_ERROR)
	}
	rhs := CreateEvaluator(pipeTarget[1])
	return NewSExpr(sym, SetLhs(lhs), SetRhs(rhs))
}

func createIf(children []parsec.Queryable) Evaler {
	condExp := CreateEvaluator(children[2])
	lhsAction := CreateEvaluator(children[3])
	rhsAction := CreateEvaluator(children[4])
	return If(condExp, lhsAction, rhsAction)
}

func createQuote(children []parsec.Queryable) Evaler {
	return Quote(CreateEvaluator(children[2]))
}
