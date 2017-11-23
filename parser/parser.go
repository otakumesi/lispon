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
	function := ast.OrdChoice("function", nil, operator, parsec.Ident())
	item := ast.OrdChoice(
		"item",
		nil,
		parsec.Int(),
		parsec.Float(),
		parsec.String(),
		quoteSymbol,
		&sexpr,
	)

	var items parsec.Parser
	items = ast.And("items", nil, item, ast.Maybe("args", nil, &items))
	sexpr = ast.And("sexpr", nil, openSexpr, function, item, items, closeSexpr)
	s := parsec.NewScanner([]byte(input))
	node, s := ast.Parsewith(sexpr, s)
	return node
}

func SexprWalk(ast parsec.Queryable) lisp.Evaluable {
	children := ast.GetChildren()
	switch ast.GetName() {
	case "sexpr":
		return createSexpr(children)
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
		symName := lisp.String(ast.GetValue())
		return lisp.Symbol{symName, true}
	}
	return lisp.Nil{}
}

func createItems(children []parsec.Queryable) lisp.Evaluable {
	if len(children) < 2 {
		return SexprWalk(children[0])
	}
	return lisp.Cons{SexprWalk(children[0]), SexprWalk(children[1])}
}

func createItem(children []parsec.Queryable) lisp.Evaluable {
	return SexprWalk(children[0])
}

func createSexpr(children []parsec.Queryable) lisp.SExpr {
	sym := lisp.Symbol{lisp.String(children[1].GetValue()), false}

	lhs := SexprWalk(children[2])
	rhs := SexprWalk(children[3])

	return lisp.NewSExpr(sym, lhs, rhs)
}
