package lexer

import (
	"fmt"
	"testing"

	"github.com/amirgamil/autumn/token"
)



func TestNextToken(t testing.T) {
	input := "=+(){},;"

	testTokens := []struct {
		expectedType token.TokenType
		expectedLiteral string 
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"}
	}

	l := New(input)
	for i, tt := range tests {
		currToken := l.NextToken()

		if currToken.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
						i, tt.expectedType, currToken.Type)
		}

		if currToken.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
						i, tt.expectedLiteral, currToken.Literal)
		}
	}

}