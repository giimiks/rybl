package main

import "rybl/lexer"

func main() {
	var tokens = lexer.StrToTokens("var const func fun while debil \"+ - +- čurák return for \\ e'ěbš ene . , = ==")
	for i := range tokens {
		println(tokens[i].Type, tokens[i].Literal)
	}
}
