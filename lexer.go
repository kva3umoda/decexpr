package decexpr

import (
	"math"

	pkgErrors "github.com/pkg/errors"
)

const (
	defaultIdentSize  = 30
	defaultNumberSize = 10
)

type Lexer struct {
	input     string
	position  int  // current position in input (points to current char)
	ch        byte // current char under examination
	prevToken Token
}

func NewLexer(input string) (*Lexer, error) {
	if len(input) >= math.MaxInt16 {
		return nil, pkgErrors.Errorf("input string too long, must be less than %d", math.MaxInt16)
	}

	l := &Lexer{
		input: input,
		prevToken: Token{
			Type:    TokenEOF,
			Literal: "",
		},
		position: -1,
	}

	l.nextChar()

	return l, nil
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	//charType :=
	switch charMap[l.ch] {
	case CharOperator:
		if l.ch == '-' {
			prevTokenType := l.prevToken.Type
			if prevTokenType == TokenEOF ||
				prevTokenType == TokenLeftParen ||
				prevTokenType == TokenComma ||
				prevTokenType == TokenOperator {
				tok = newToken(TokenUnaryOperator, string(l.ch), l.position)
			} else {
				tok = newToken(TokenOperator, string(l.ch), l.position)
			}

			break
		}

		tok = newToken(TokenOperator, string(l.ch), l.position)
	case CharLeftParen:
		tok = newToken(TokenLeftParen, string(l.ch), l.position)
	case CharRightParen:
		tok = newToken(TokenRightParen, string(l.ch), l.position)
	case CharComma:
		tok = newToken(TokenComma, string(l.ch), l.position)
	case CharDigit:
		tok = l.readNumber()

		l.prevToken = tok

		return tok
	case CharLetter:
		tok = l.readIdentifier()
		if l.ch == '(' {
			tok.Type = TokenFunction
		}

		l.prevToken = tok

		return tok
	case CharEOF:
		tok = newToken(TokenEOF, "", l.position)
	default:
		tok = newToken(TokenIllegal, string(l.ch), l.position)
	}

	l.nextChar()

	l.prevToken = tok

	return tok
}

func (l *Lexer) Tokenize() ([]Token, error) {
	tokens := make([]Token, 0, 10)

	for token := l.NextToken(); token.Type != TokenEOF; token = l.NextToken() {
		if token.Type == TokenIllegal {
			return nil, pkgErrors.Errorf("invalid token type: %s, position:%d", token.Literal, l.position)
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (l *Lexer) nextChar() {
	l.position++

	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
}

func (l *Lexer) readIdentifier() Token {
	beginPos := l.position
	buf := make([]byte, 0, defaultIdentSize)

	for isLetter(l.ch) || isDigit(l.ch) {
		buf = append(buf, l.ch)

		l.nextChar()
	}

	return newToken(TokenIdent, string(buf), beginPos)
}

func (l *Lexer) readNumber() Token {
	beginPos := l.position
	exp := int16(0)
	///decimalFound := false

	digits := make([]byte, 0, defaultNumberSize)
	tokenType := TokenIntNumber

	for isDigit(l.ch) || isDot(l.ch) {
		if l.ch == '.' {
			if tokenType == TokenFloatNumber {
				break
			}

			tokenType = TokenFloatNumber
			l.nextChar()

			continue
		}

		digits = append(digits, l.ch)

		if tokenType == TokenFloatNumber {
			exp++
		}

		l.nextChar()
	}

	tok := newToken(TokenFloatNumber, string(digits), beginPos)
	tok.Exp = exp

	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.nextChar()
	}
}
