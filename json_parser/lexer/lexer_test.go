package lexer

import (
	"testing"
)

func TestSimpleValidJson(t *testing.T) {
	lexer := NewLexer("{}")
	expectedTokens := []Token{
		{LBRACE, "{"},
		{RBRACE, "}"},
	}

	for _, expectedToken := range expectedTokens {
		lexerToken := lexer.NextToken()

		if lexerToken.Type != expectedToken.Type || lexerToken.Literal != expectedToken.Literal {
			t.Fatalf("Expected %q got %q", expectedToken, lexerToken)
		}
	}
}

func TestSimpleInvalidJson(t *testing.T) {
	lexer := NewLexer("{.}")
	expectedTokens := []Token{
		{LBRACE, "{"},
		{ILLEGAL, "."},
		{RBRACE, "}"},
	}

	for _, expectedToken := range expectedTokens {
		lexerToken := lexer.NextToken()

		if lexerToken.Type != expectedToken.Type || lexerToken.Literal != expectedToken.Literal {
			t.Fatalf("Expected %q got %q", expectedToken, lexerToken)
		}
	}
}

func TestNumberValidJson(t *testing.T) {
	lexer := NewLexer("{ \"number\": 10.2, \"location\": \"Portugal\" }")
	expectedTokens := []Token{
		{LBRACE, "{"},
		{STRING, "\"number\""},
		{COLON, ":"},
		{NUMBER, "10.2"},
		{COMMA, ","},
		{STRING, "\"location\""},
		{COLON, ":"},
		{STRING, "\"Portugal\""},
		{RBRACE, "}"},
	}

	for _, expectedToken := range expectedTokens {
		lexerToken := lexer.NextToken()

		if lexerToken.Type != expectedToken.Type || lexerToken.Literal != expectedToken.Literal {
			t.Fatalf("Expected %q got %q", expectedToken, lexerToken)
		}
	}
}
