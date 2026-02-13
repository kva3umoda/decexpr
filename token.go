package decexpr

type TokenType byte

const (
	TokenIllegal TokenType = iota + 1
	TokenEOF
	TokenNumber
	TokenOperator
	TokenUnaryOperator
	TokenIdent
	TokenFunction
	TokenLeftParen
	TokenRightParen
	TokenComma
)

func (tt TokenType) String() string {
	switch tt {
	case TokenIllegal:
		return "Illegal"
	case TokenEOF:
		return "EOF"
	case TokenNumber:
		return "Number"
	case TokenUnaryOperator:
		return "UnaryOperator"
	case TokenOperator:
		return "Operator"
	case TokenIdent:
		return "Ident"
	case TokenFunction:
		return "Function"
	case TokenLeftParen:
		return "LeftParen"
	case TokenRightParen:
		return "RightParen"
	case TokenComma:
		return "Comma"
	default:
		return "Unknown"
	}
}

type Token struct {
	Type     TokenType
	Exp      int16
	Position int16
	Literal  string
}

func newToken(tokenType TokenType, literal string, position int) Token {
	return Token{
		Type:     tokenType,
		Literal:  literal,
		Position: int16(position),
	}
}
