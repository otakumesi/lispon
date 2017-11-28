package lisp

import (
	"github.com/prataprc/goparsec"
)

func Parse(input string) parsec.Queryable {
	var expr parsec.Parser
	var lambdaExpr parsec.Parser
	var pipeExpr parsec.Parser
	var isAtomExpr parsec.Parser
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
		&pipeExpr,
		&isAtomExpr,
		&expr,
	)

	var items parsec.Parser
	items = ast.And("items", nil, item, ast.Maybe("args", nil, &items))

	operator := parsec.Token(`[+-/*%]`, "OPERATOR")
	symbol := ast.OrdChoice("symbol", nil, operator, parsec.Ident())
	maybeItems := ast.Maybe("maybeItems", nil, items)

	expr = ast.And(
		"expr",
		nil,
		openSexpr,
		symbol,
		ast.Maybe("lhs", nil, item),
		maybeItems,
		closeSexpr,
	)

	lambda := parsec.Token(`lambda|->`, "LAMBDA")
	lambdaExpr = ast.And(
		"lambdaExpr",
		nil,
		openSexpr,
		lambda,
		openSexpr,
		maybeItems,
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

	pipe := parsec.Token(`\|>`, "PIPE")
	pipeTarget := ast.And(
		"pipeTarget",
		nil,
		symbol,
		maybeItems,
	)
	pipeExpr = ast.And(
		"pipeExpr",
		nil,
		openSexpr,
		item,
		pipe,
		pipeTarget,
		closeSexpr,
	)

	isAtom := parsec.Token(`atom`, "IS_ATOM")
	isAtomExpr = ast.And(
		"isAtomExpr",
		nil,
		openSexpr,
		isAtom,
		items,
		closeSexpr,
	)

	sexpr := ast.OrdChoice("sexpr", nil, defineExpr, isAtomExpr, pipeExpr, lambdaExpr, expr)

	s := parsec.NewScanner([]byte(input))
	node, s := ast.Parsewith(sexpr, s)

	return node
}
