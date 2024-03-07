package main

import (
	"github.com/henrmota/codechallenges/json_parser/lexer"
	"github.com/henrmota/codechallenges/json_parser/parser"
)

func main() {
	lexer := lexer.NewLexer("{\"Henrique\": \"Mota\", \"stuff\": { \"age\": \"40\"} }")
	parser := parser.NewParser(lexer)

	parser.Parse()
}
