package lisp

import (
	"github.com/prataprc/goparsec"
)

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

	lambda := parsec.Token(`lambda|->`, "LAMBDA")
	lambdaExpr = ast.And(
		"lambdaExpr",
		nil,
		openSexpr,
		lambda,
		openSexpr,
		ast.Maybe("lambdaArgs", nil, items),
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
