package decexpr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	exp := "5 + 10 -func(test,123)"

	lex, err := NewLexer(exp)
	assert.Nil(t, err)

	tokens, err := lex.Tokenize()
	assert.NoError(t, err)

	fmt.Println(tokens)
}

func TestLexer_Tokenize(t *testing.T) {
	tests := []struct {
		expr   string
		tokens []Token
	}{
		{
			expr: "5+10",
			tokens: []Token{
				{Type: TokenNumber, Literal: "5", Exp: 0, Position: 0},
				{Type: TokenOperator, Literal: "+", Exp: 0, Position: 1},
				{Type: TokenNumber, Literal: "10", Exp: 0, Position: 2},
			},
		},
		{
			expr: "var1 + var2(1,432.5)",
			tokens: []Token{
				{Type: TokenIdent, Literal: "var1", Exp: 0, Position: 0},
				{Type: TokenOperator, Literal: "+", Exp: 0, Position: 5},
				{Type: TokenFunction, Literal: "var2", Exp: 0, Position: 7},
				{Type: TokenLeftParen, Literal: "(", Exp: 0, Position: 11},
				{Type: TokenNumber, Literal: "1", Exp: 0, Position: 12},
				{Type: TokenComma, Literal: ",", Exp: 0, Position: 13},
				{Type: TokenNumber, Literal: "4325", Exp: 1, Position: 14},
				{Type: TokenRightParen, Literal: ")", Exp: 0, Position: 19},
			},
		},
		{
			expr: "5 + 10",
			tokens: []Token{
				{Type: TokenNumber, Literal: "5", Exp: 0, Position: 0},
				{Type: TokenOperator, Literal: "+", Exp: 0, Position: 2},
				{Type: TokenNumber, Literal: "10", Exp: 0, Position: 4},
			},
		},
		{
			expr: "5-10",
			tokens: []Token{
				{Type: TokenNumber, Literal: "5", Exp: 0, Position: 0},
				{Type: TokenOperator, Literal: "-", Exp: 0, Position: 1},
				{Type: TokenNumber, Literal: "10", Exp: 0, Position: 2},
			},
		},
		{
			expr: "5.10",
			tokens: []Token{
				{Type: TokenNumber, Literal: "510", Exp: 2, Position: 0},
			},
		},
		{
			expr: "-5.10",
			tokens: []Token{
				{Type: TokenUnaryOperator, Literal: "-", Exp: 0, Position: 0},
				{Type: TokenNumber, Literal: "510", Exp: 2, Position: 1},
			},
		},
		{
			expr: "-var1  * (-10+log(var2, 10.123)-max(-45.56))",
			tokens: []Token{
				{Type: TokenUnaryOperator, Literal: "-", Exp: 0, Position: 0},
				{Type: TokenIdent, Literal: "var1", Exp: 0, Position: 1},
				{Type: TokenOperator, Literal: "*", Exp: 0, Position: 7},
				{Type: TokenLeftParen, Literal: "(", Exp: 0, Position: 9},
				{Type: TokenUnaryOperator, Literal: "-", Exp: 0, Position: 10},
				{Type: TokenNumber, Literal: "10", Exp: 0, Position: 11},
				{Type: TokenOperator, Literal: "+", Exp: 0, Position: 13},
				{Type: TokenFunction, Literal: "log", Exp: 0, Position: 14},
				{Type: TokenLeftParen, Literal: "(", Exp: 0, Position: 17},
				{Type: TokenIdent, Literal: "var2", Exp: 0, Position: 18},
				{Type: TokenComma, Literal: ",", Exp: 0, Position: 22},
				{Type: TokenNumber, Literal: "10123", Exp: 3, Position: 24},
				{Type: TokenRightParen, Literal: ")", Exp: 0, Position: 30},
				{Type: TokenOperator, Literal: "-", Exp: 0, Position: 31},
				{Type: TokenFunction, Literal: "max", Exp: 0, Position: 32},
				{Type: TokenLeftParen, Literal: "(", Exp: 0, Position: 35},
				{Type: TokenUnaryOperator, Literal: "-", Exp: 0, Position: 36},
				{Type: TokenNumber, Literal: "4556", Exp: 2, Position: 37},
				{Type: TokenRightParen, Literal: ")", Exp: 0, Position: 42},
				{Type: TokenRightParen, Literal: ")", Exp: 0, Position: 43},
			},
		},
		{
			expr: "5 + -10",
			tokens: []Token{
				{Type: TokenNumber, Literal: "5", Exp: 0, Position: 0},
				{Type: TokenOperator, Literal: "+", Exp: 0, Position: 2},
				{Type: TokenUnaryOperator, Literal: "-", Exp: 0, Position: 4},
				{Type: TokenNumber, Literal: "10", Exp: 0, Position: 5},
			},
		},
		{
			expr: "5+var1",
			tokens: []Token{
				{Type: TokenNumber, Literal: "5", Exp: 0, Position: 0},
				{Type: TokenOperator, Literal: "+", Exp: 0, Position: 1},
				{Type: TokenIdent, Literal: "var1", Exp: 0, Position: 2},
			},
		},
		{
			expr: "3 + 4 * 2 / (1 - 5)^2",
			tokens: []Token{
				{Type: TokenNumber, Literal: "3", Exp: 0, Position: 0},
				{Type: TokenOperator, Literal: "+", Exp: 0, Position: 2},
				{Type: TokenNumber, Literal: "4", Exp: 0, Position: 4},
				{Type: TokenOperator, Literal: "*", Exp: 0, Position: 6},
				{Type: TokenNumber, Literal: "2", Exp: 0, Position: 8},
				{Type: TokenOperator, Literal: "/", Exp: 0, Position: 10},
				{Type: TokenLeftParen, Literal: "(", Exp: 0, Position: 12},
				{Type: TokenNumber, Literal: "1", Exp: 0, Position: 13},
				{Type: TokenOperator, Literal: "-", Exp: 0, Position: 15},
				{Type: TokenNumber, Literal: "5", Exp: 0, Position: 17},
				{Type: TokenRightParen, Literal: ")", Exp: 0, Position: 18},
				{Type: TokenOperator, Literal: "^", Exp: 0, Position: 19},
				{Type: TokenNumber, Literal: "2", Exp: 0, Position: 20},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			lex, err := NewLexer(test.expr)
			assert.NoError(t, err)
			tokens, err := lex.Tokenize()
			assert.NoError(t, err)
			assert.EqualValues(t, test.tokens, tokens)
		})
	}
}

// BenchmarkLexer_Tokenize-12    	 1189868	       973.1 ns/op	    1140 B/op	      16 allocs/op
// BenchmarkLexer_Tokenize-12    	 1341477	       897.0 ns/op	     885 B/op	      24 allocs/op
// BenchmarkLexer_Tokenize-12    	 1424876	       825.0 ns/op	     869 B/op	      23 allocs/op
// BenchmarkLexer_Tokenize-12    	 1473672	       813.8 ns/op	     869 B/op	      23 allocs/op
func BenchmarkLexer_Tokenize(b *testing.B) {
	exp := "-var1  * (-10+log(var2, 10.123)-max(-45.56))"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l, err := NewLexer(exp)
		if err != nil {
			b.Fatal(err)
		}

		lst, err := l.Tokenize()
		if err != nil {
			b.Fatal(err)
		}

		lst = lst
	}
}
