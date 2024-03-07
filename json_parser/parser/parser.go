package parser

import (
	"fmt"
	"log"

	"github.com/henrmota/codechallenges/json_parser/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer,
	}
}

func (parser Parser) Parse() {
	token := parser.lexer.NextToken()

	switch token.Type {
	case lexer.LBRACE:
		parser.parseObject()
	case lexer.RARRAY:
		parser.parseArray()
	default:
		log.Fatalf("Unexpect character at the begining %s\n", token.Literal)
	}
}

func (parser Parser) parseObject() {
	nextToken := parser.lexer.NextToken()
	if nextToken.Type == lexer.RBRACE {
		return
	}

	for {
		if nextToken.Type != lexer.STRING {
			log.Fatalf("Expecting a string as a key")
		}

		nextToken := parser.lexer.NextToken()
		if nextToken.Type != lexer.COLON {
			log.Fatalf("Expecting a colon after the key")
		}

		parser.parseValue()

		nextToken = parser.lexer.NextToken()
		if nextToken.Type == lexer.RBRACE {
			return
		}

		if nextToken.Type != lexer.COMMA {
			log.Fatalf("Expecting a comma")
		}

		nextToken = parser.lexer.NextToken()
	}
}

func (parser *Parser) parseValue() {
	nextToken := parser.lexer.NextToken()
	if nextToken.Type == lexer.LBRACE {
		parser.parseObject()
		return
	}

	if nextToken.Type == lexer.LARRAY {
		parser.parseArray()
		return
	}

	if nextToken.Type != lexer.STRING {
		log.Fatalf("Expecting a string as a value")
	}
}

func (parser *Parser) parseArray() {
	fmt.Printf("Array")
}
