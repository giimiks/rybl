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
	"number": true,
	"string": true,
	"bool": true,
	"array": true,
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
	Row int
	Col int
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '"'
}

func isNumber(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isDelimiter(ch rune) bool {
	switch ch {
	case '(', ')', '{', '}', ';', ',', '.':
		return true
	default:
		return false
	}
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
	t.Type = StringLiteral
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

func isEOL(ch rune) bool {
	return ch == '\n'
}

/*
TODO:
	Dont ignore dots and commas and semicolons and shit
*/
func StrToTokens(str string) (tokens []Token) {
	var i = 0
	var row, col = 0,0
	for i < len(str) {
		r, size := utf8.DecodeRuneInString(str[i:])
		switch {
		case isLetter(r):
			var token, index = buildFromLetters(i, str)
			token.Col = col
			token.Row = row
			tokens = append(tokens, token)
			i = index
		case isEOL(r):
			col = 0
			row+=1
			i+=size
		case isNumber(r):
			var token, index = buildNumeric(i, str)
			token.Col = col
			token.Row = row
			tokens = append(tokens, token)
			i = index
		case isOp(r):
			var token = Token{Type: Operator, Literal: string(r), Row: row, Col: col}	
			tokens = append(tokens, token)
			i+=size
		case isWhSpace(r):
			col = i
			i+=size
		case isDelimiter(r):
			var token = Token{Type: Delimiter, Literal: string(r), Row: row, Col: col}	
			tokens = append(tokens, token)
			i+=size
		default:
			i+=size

		}
		col++
	}
	return tokens
}
