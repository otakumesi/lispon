package parser

import (
	"strconv"

	"../lisp"
	"github.com/prataprc/goparsec"
)

var RUNTIME_ERROR = "Runtime Error"

func Parse(input string) parsec.Queryable {
	var expr parsec.Parser
	var lambdaExpr parsec.Parser
	ast := parsec.NewAST("lisp", 100)

	openSexpr := parsec.Atom("(", "OPEN_SEXPR")
	closeSexpr := parsec.Atom(")", "CLOSE_SEXPR")
	quoteSymbol := parsec.Token(`'[A-Za-z][0-9a-zA-Z_]*`, "QUOTED_SYMBOL")
	string := parsec.Token(`".*?"`, "STRING")

	item := ast.OrdChoice(
		"item",
		nil,
		parsec.Int(),
		parsec.Float(),
		parsec.Ident(),
		quoteSymbol,
		string,
		&lambdaExpr,
		&expr,
	)

	var items parsec.Parser
	items = ast.And("items", nil, item, ast.Maybe("args", nil, &items))

	operator := parsec.Token(`[+-/*%]`, "OPERATOR")
	symbol := ast.OrdChoice("symbol", nil, operator, parsec.Ident())
	expr = ast.And(
		"expr",
		nil,
		openSexpr,
		symbol,
		ast.Maybe("lhs", nil, item),
		ast.Maybe("rhs", nil, items),
		closeSexpr,
	)

	lambda := parsec.Token(`lambda`, "LAMBDA")
	lambdaExpr = ast.And(
		"lambdaExpr",
		nil,
		openSexpr,
		lambda,
		openSexpr,
		items,
		closeSexpr,
		expr,
		closeSexpr,
	)

	define := parsec.Token(`define`, "DEFINE")
	defineExpr := ast.And(
		"defineExpr",
		nil,
		openSexpr,
		define,
		quoteSymbol,
		items,
		closeSexpr,
	)

	sexpr := ast.OrdChoice("sexpr", nil, defineExpr, lambdaExpr, expr)

	s := parsec.NewScanner([]byte(input))
	node, s := ast.Parsewith(sexpr, s)
	return node
}

func ParseSExpr(ast parsec.Queryable) lisp.Evaluable {
	children := ast.GetChildren()
	switch ast.GetName() {
	case "sexpr":
		return ParseSExpr(ast.GetChildren()[0])
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
		return lisp.Number(i)
	case "FLOAT":
		f, err := strconv.ParseFloat(ast.GetValue(), 64)
		if err != nil {
			panic(err)
		}
		return lisp.Number(f)
	case "STRING":
		rawStr := ast.GetValue()
		return lisp.String(rawStr[1 : len(rawStr)-1])
	case "IDENT":
		return lisp.NewSymbol(ast.GetValue())
	case "QUOTED_SYMBOL":
		rawStr := ast.GetValue()
		return lisp.NewSymbol(rawStr[1:len(rawStr)], lisp.SetIsQuote(true))
	}
	return lisp.Nil{}
}

func createItems(children []parsec.Queryable) lisp.Evaluable {
	if children[1].GetName() == "missing" {
		return ParseSExpr(children[0])
	}
	return lisp.Cons{ParseSExpr(children[0]), ParseSExpr(children[1])}
}

func createItem(children []parsec.Queryable) lisp.Evaluable {
	return ParseSExpr(children[0])
}

func createExpr(children []parsec.Queryable) lisp.SExpr {
	sym := lisp.NewSymbol(children[1].GetValue())

	if len(children) < 2 {
		return lisp.NewSExpr(sym)
	}
	lhs := ParseSExpr(children[2])

	if len(children) < 3 {
		return lisp.NewSExpr(sym, lisp.SetLhs(lhs))
	}
	rhs := ParseSExpr(children[3])

	return lisp.NewSExpr(sym, lisp.SetLhs(lhs), lisp.SetRhs(rhs))
}

func createDefine(children []parsec.Queryable) lisp.Evaluable {
	argSym, ok := ParseSExpr(children[2]).(lisp.Symbol)

	if !ok {
		panic(RUNTIME_ERROR)
	}

	expr := ParseSExpr(children[3])
	return lisp.Define(argSym, expr)
}

func createLambda(children []parsec.Queryable) lisp.Evaluable {
	form, ok := ParseSExpr(children[5]).(lisp.SExpr)

	if !ok {
		panic(RUNTIME_ERROR)
	}

	var args []lisp.Symbol
	for _, child := range children[3].GetChildren() {
		argSym, ok := ParseSExpr(child).(lisp.Symbol)
		if !ok {
			panic(RUNTIME_ERROR)
		}
		args = append(args, argSym)
	}
	return lisp.Lambda(form, args...)
}
