package lexer

import (
	"unicode"
	"unicode/utf8"
)

type TokenType string

const (
	Keyword       TokenType = "Keyword"
	Identifier    TokenType = "Identifier"
	StringLiteral TokenType = "StringLiteral"
	NumberLiteral TokenType = "NumberLiteral"
	Operator      TokenType = "Operator"
	Delimiter     TokenType = "Delimiter"
)

type Token struct {
	Type    TokenType
	Literal string
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}

func isNumber(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isOp(ch rune) bool {
	switch {
	case ch == '+':
		return true
	case ch == '-':
		return true
	case ch == '=':
		return true
	case ch == '*':
		return true
	case ch == '/':
		return true
	default:
		return false
	}
}

func isWhSpace(ch rune) bool {
	return unicode.IsSpace(ch) && ch != '\n'
}

func buildNumeric(i int, str string) (Token, int) {
	var t = Token{}
	t.Type = NumberLiteral
	var literal []rune
	for i < len(str) {
		r, size := utf8.DecodeRuneInString(str[i:])
		if !isNumber(r) {
			break
		}
		literal = append(literal, r)
		i+=size
	}
	t.Literal = string(literal)
	return t,i
}

func strToTokens(str string) (tokens []Token) {
	for _, ch := range str {
		switch {
		case isLetter(ch):
			return
		case isNumber(ch):
			return
		case isOp(ch):
			return
		case isWhSpace(ch):
			return
		default:
			return
		}
	}
	return
}
