package lexer

import (
	"fmt"
	"log"
)

type TokenType int

const (
	LBRACE TokenType = iota
	RBRACE
	LARRAY
	RARRAY
	COLON
	COMMA
	STRING
	NUMBER
	LKEY
	LVALUE
	ILLEGAL
	EOF
)

type Lexer struct {
	input        string
	char         byte
	position     int
	readPosition int
	Line         int
	Column       int
}

type Token struct {
	Type    TokenType
	Literal string
}

func NewLexer(contents string) *Lexer {
	currentChar := contents[0]

	l := &Lexer{contents, currentChar, 0, 0, 0, 0}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipSpace()

	switch l.char {
	case '{':
		tok = Token{LBRACE, "{"}
	case '}':
		tok = Token{RBRACE, "}"}
	case '[':
		tok = Token{LARRAY, "["}
	case ']':
		tok = Token{RARRAY, "]"}
	case ':':
		tok = Token{COLON, ":"}
	case ',':
		tok = Token{COMMA, ","}
	case '"':
		literal, err := l.getString()
		if err != nil {
			tok.Literal = err.Error()
			tok.Type = ILLEGAL

			return tok
		}

		tok = Token{STRING, literal}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isNumeric(l.char) {
			var err error
			tok.Literal, err = l.readNumber()
			tok.Type = NUMBER
			if err != nil {
				log.Fatal(err)
			}
			return tok
		}

		tok.Literal = string(l.char)
		tok.Type = ILLEGAL
	}

	l.readChar()

	return tok
}

func (l *Lexer) readNumber() (string, error) {
	literal := []byte{l.char}
	l.readChar()
	for l.char >= '0' && l.char <= '9' {
		literal = append(literal, l.char)
		l.readChar()
	}

	if l.char == '.' {
		literal = append(literal, l.char)
	}

	l.readChar()
	for l.char >= '0' && l.char <= '9' {
		literal = append(literal, l.char)
		l.readChar()
	}

	if l.char == 0 {
		return "", fmt.Errorf("unexpected End of File")
	}

	return string(literal), nil
}

func (l *Lexer) getString() (string, error) {
	literal := []byte{l.char}
	l.readChar()
	for l.char != '"' && l.char != 0 {
		literal = append(literal, l.char)
		l.readChar()
	}

	literal = append(literal, l.char)

	if l.char == 0 {
		return "", fmt.Errorf("unexpected End of File")
	}

	return string(literal), nil
}

func (l *Lexer) skipSpace() {
	if l.char == '\n' {
		l.Line += 1
		l.Column = 0
	}

	for l.char == ' ' || l.char == '\n' || l.char == '\t' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	l.Column += 1
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) PeekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func isNumeric(c byte) bool {
	return c >= '0' && c <= '9' || c == '-'
}
