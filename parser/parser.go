package parser

import (
	"strconv"

	"../lisp"
	"github.com/prataprc/goparsec"
)

func Parse(input string) parsec.Queryable {
	var sexpr parsec.Parser
	ast := parsec.NewAST("lisp", 100)

	openSexpr := parsec.Atom("(", "OPEN_SEXPR")
	closeSexpr := parsec.Atom(")", "CLOSE_SEXPR")
	operator := parsec.Token(`[+-/*%]`, "OPERATOR")
	quoteSymbol := parsec.Token(`'[A-Za-z][0-9a-zA-Z_]*`, "QUOTED_SYMBOL")
	symbol := ast.OrdChoice("symbol", nil, operator, parsec.Ident())
	item := ast.OrdChoice(
		"item",
		nil,
		parsec.Int(),
		parsec.Float(),
		parsec.String(),
		parsec.Ident(),
		quoteSymbol,
		&sexpr,
	)

	var items parsec.Parser
	items = ast.And("items", nil, item, ast.Maybe("args", nil, &items))
	sexpr = ast.And(
		"sexpr",
		nil,
		openSexpr,
		symbol,
		ast.Maybe("lhs", nil, item),
		ast.Maybe("rhs", nil, items),
		closeSexpr,
	)

	s := parsec.NewScanner([]byte(input))
	node, s := ast.Parsewith(sexpr, s)
	return node
}

func ParseSExpr(ast parsec.Queryable) lisp.Evaluable {
	children := ast.GetChildren()
	switch ast.GetName() {
	case "sexpr":
		return createSExpr(children)
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
		return lisp.String(ast.GetValue())
	case "IDENT":
		symName := lisp.String(ast.GetValue())
		return lisp.Symbol{symName, false}
	case "QUOTED_SYMBOL":
		symbol := ast.GetValue()
		symName := lisp.String(symbol[1:len(symbol)])
		return lisp.Symbol{symName, true}
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

func createSExpr(children []parsec.Queryable) lisp.SExpr {
	sym := lisp.Symbol{lisp.String(children[1].GetValue()), false}

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
