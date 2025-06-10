package parser

import "rybl/lexer"

type Expr interface{}

type Token lexer.Token

type BinaryExpr struct {
	Left  Expr
	Op    Token
	Right Expr
}

type LiteralExpr struct {
	Value string
}

type VarExpr struct {
	Name Token
}
