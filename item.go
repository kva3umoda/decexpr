package decexpr

import (
	"strconv"

	pkgErrors "github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

/// Оператор	Описание					Приоритет	Ассоциативность
//	()			Скобки						0			—
//	sin, max,..	Функции						5			—
//	^			Возведение в степень		4			Справа налево
//	-, +, !		Унарные операции			3			Справа налево
//	*, /, %		Умножение, деление			2			Слева направо
//	+, -		Бинарное сложение/вычитание	1			Слева направо

type RPNItem struct {
	Token
	Priority     int
	FuncArgCount int
	Number       decimal.Decimal
}

func NewRPNItem(token Token) (*RPNItem, error) {
	item := &RPNItem{
		Token: token,
	}

	switch token.Type {
	case TokenIdent, TokenEOF:
		item.Priority = 0
	case TokenFloatNumber:
		value, err := strconv.ParseInt(token.Literal, 10, 64)
		if err != nil {
			return item, err
		}

		item.Number = decimal.New(value, -int32(token.Exp))
		item.Priority = 0
	case TokenIntNumber:
		value, err := strconv.ParseInt(token.Literal, 10, 64)
		if err != nil {
			return item, err
		}

		item.Number = decimal.NewFromInt(value)
	case TokenFunction:
		item.Priority = 5
	case TokenUnaryOperator:
		item.Priority = 4
	case TokenLeftParen, TokenRightParen:
		item.Priority = 0
	case TokenOperator:
		priority, exists := operatorPriority[token.Literal]
		if !exists {
			return item, pkgErrors.Errorf("invalid operator priority: %s", token.Literal)
		}
		item.Priority = priority
	case TokenComma:
		item.Priority = 0
	default:
		panic("unhandled default case")
	}

	return item, nil
}

type Operator = string

const (
	OpAdd   Operator = "+"
	OpSub   Operator = "-"
	OpMul   Operator = "*"
	OpDiv   Operator = "/"
	OpMod   Operator = "%"
	OpPower Operator = "^"
)

var operatorPriority = map[Operator]int{
	OpAdd:   1,
	OpSub:   1,
	OpMul:   2,
	OpDiv:   2,
	OpMod:   2,
	OpPower: 4,
}
