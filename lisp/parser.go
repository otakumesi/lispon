package lisp

import (
	"github.com/prataprc/goparsec"
)

func Parse(input string) parsec.Queryable {
	var expr parsec.Parser
	var lambdaExpr parsec.Parser
	var ifExpr parsec.Parser
	var pipeExpr parsec.Parser
	var quoteExpr parsec.Parser
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
		&ifExpr,
		&quoteExpr,
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

	ifSym := parsec.Token(`if`, "IF")
	ifExpr = ast.And(
		"ifExpr",
		nil,
		openSexpr,
		ifSym,
		ast.OrdChoice("condExp", nil, expr, item),
		ast.OrdChoice("lhsAction", nil, expr, item),
		ast.OrdChoice("rhsAction", nil, expr, item),
		closeSexpr,
	)

	quoteSym := parsec.Token(`quote`, "QUOTE")
	quoteExpr = ast.And(
		"quoteExpr",
		nil,
		openSexpr,
		quoteSym,
		ast.OrdChoice("quoteValues", nil, item, items, expr),
		closeSexpr,
	)

	sexpr := ast.OrdChoice("sexpr", nil, defineExpr, pipeExpr, ifExpr, quoteExpr, lambdaExpr, expr)

	var sexprs parsec.Parser
	sexprs = ast.Many("sexprs", nil, sexpr)

	s := parsec.NewScanner([]byte(input))
	node, s := ast.Parsewith(sexprs, s)

	return node
}
