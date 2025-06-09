package lexer

import (
	"unicode"
	"unicode/utf8"
)

var keyword_set = map[string]bool {
	"if": true,
	"else": true,
	"return": true,
	"proc": true,
	"var": true,
	"const": true,
	"for" : true,
	"while" : true,
}

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
		var r, size = utf8.DecodeRuneInString(str[i:])
		if !isNumber(r) {
			break
		}
		literal = append(literal, r)
		i+=size
	}
	t.Literal = string(literal)
	return t,i
}

func buildStrLiteral(i int, str string, t Token) (Token, int) {
	var literal []rune
	for i < len(str) {
		var r, size =  utf8.DecodeRuneInString(str[i:])
		if r == '"' {
			i+=size
			break
		}
		literal = append(literal, r)
		i += size
	}
	t.Literal = string(literal)
	return t, i

}

func buildIdentOrKw(i int, str string, t Token) (Token, int) {
	var literal []rune
	for i < len(str) {
		var r, size =  utf8.DecodeRuneInString(str[i:])

		//shouldnt really ignore non-letters, but fine for now
		if !isLetter(r) {
			break
		}
		literal = append(literal, r)
		i += size
	}
	t.Literal = string(literal)
	determineIdentOrKw(&t)
	return t,i
}

func buildFromLetters(i int, str string) (Token, int) {
	var t = Token{}
	var r, _ =  utf8.DecodeRuneInString(str[i:])
	if r == '"' {
		t.Type = StringLiteral
		return buildStrLiteral(i+1, str, t)
	} else {
		return buildIdentOrKw(i, str, t)
	}

}

func determineIdentOrKw(t *Token) {
	if keyword_set[t.Literal] {
		t.Type = Keyword
	} else {
		t.Type = Identifier
	}
}

/*
TODO:
	Dont ignore dots and commas and semicolons and shit
*/
func strToTokens(str string) (tokens []Token) {
	var i = 0
	for i < len(str) {
		r, size := utf8.DecodeRuneInString(str[i:])
		switch {
		case isLetter(r):
			var token, index = buildFromLetters(i, str)
			tokens = append(tokens, token)
			i = index
		case isNumber(r):
			var token, index = buildNumeric(i, str)
			tokens = append(tokens, token)
			i = index
		case isOp(r):
			tokens = append(tokens, Token{Type: Operator, Literal: string(r)})
			i+=size
		case isWhSpace(r):
			i+=size
		default:
			i+=size
		}
	}
	return
}
