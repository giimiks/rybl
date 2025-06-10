package main

import (
	"fmt"
	"rybl/lexer"
)

func main() {
	var tokens = lexer.StrToTokens("var promenna = 512\n;while (promenna > 0) {\npromenna-=1\n		}")
	for i := range tokens {
		fmt.Printf("Type: %s Val: %s Row,Col: %d,%d\n",tokens[i].Type,tokens[i].Literal, tokens[i].Row, tokens[i].Col)
	}
}
