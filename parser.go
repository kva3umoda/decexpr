package decexpr

import (
	pkgErrors "github.com/pkg/errors"
)

type Parser struct {
	functions map[string]int
}

func NewParser(functions map[string]int) *Parser {
	return &Parser{
		functions: functions,
	}
}

func (p *Parser) Parse(exp string) ([]*RPNItem, error) {
	l, err := NewLexer(exp)
	if err != nil {
		return nil, err
	}

	output := make([]*RPNItem, 0, len(exp))
	itemStack := NewStackItems()
	argsSkack := NewArgStack()

	for token := l.NextToken(); token.Type != TokenEOF; token = l.NextToken() {
		if token.Type == TokenIllegal {
			return nil, pkgErrors.Errorf("invalid token type: %s, posistion:%d",
				token.Literal, token.Position)
		}

		newItem, err := NewRPNItem(token)
		if err != nil {
			return nil, pkgErrors.Wrapf(err, "invalid number token: %s-%s, posistion:%d",
				token.Literal, token.Type.String(), token.Position)
		}

		switch token.Type {
		case TokenFloatNumber, TokenIdent:
			output = append(output, newItem)
		case TokenOperator, TokenUnaryOperator:
			for itemStack.Len() > 0 {
				item := itemStack.Pop()

				if item.Priority < newItem.Priority {
					itemStack.Push(item)

					break
				}

				if item.Type == TokenFunction {
					item.FuncArgCount = argsSkack.Pop()

					if err := p.checkFunction(item); err != nil {
						return nil, err
					}
				}

				output = append(output, item)
			}

			itemStack.Push(newItem)
		case TokenLeftParen:
			itemStack.Push(newItem)
		case TokenRightParen:
			closed := false
			for itemStack.Len() > 0 {
				item := itemStack.Pop()
				if item.Type == TokenLeftParen {
					closed = true

					break
				} else if item.Type == TokenFunction {
					item.FuncArgCount = argsSkack.Pop()

					if err := p.checkFunction(item); err != nil {
						return nil, err
					}
				}

				output = append(output, item)
			}

			if !closed {
				return nil, pkgErrors.Errorf("invalid left paren token: %s-%s, posistion:%d",
					token.Literal, token.Type.String(), token.Position)
			}
		case TokenFunction:
			itemStack.Push(newItem)
			argsSkack.Push()
		case TokenComma:
			for itemStack.Len() > 0 {
				item := itemStack.Pop()
				if item.Type == TokenLeftParen {
					itemStack.Push(item)

					break
				} else if item.Type == TokenFunction {
					item.FuncArgCount = argsSkack.Pop()

					if err := p.checkFunction(item); err != nil {
						return nil, err
					}
				}

				output = append(output, item)
			}

			argsSkack.Inc()
		default:
			return nil, pkgErrors.Wrapf(err, "invalid number token: %s-%s, posistion:%d",
				token.Literal, token.Type.String(), l.position)
		}
	}

	for itemStack.Len() > 0 {
		item := itemStack.Pop()

		switch item.Type {
		case TokenLeftParen:
			return nil, pkgErrors.Errorf("invalid left paren token, posistion:%d", item.Position)
		case TokenRightParen:
			return nil, pkgErrors.Errorf("invalid right paren token, posistion:%d", item.Position)
		case TokenFunction:
			item.FuncArgCount = argsSkack.Pop()

			if err := p.checkFunction(item); err != nil {
				return nil, err
			}
		default:
			break
		}

		output = append(output, item)
	}

	return output, nil
}

func (p *Parser) checkFunction(item *RPNItem) error {
	funcArgs, exists := p.functions[item.Literal]
	if !exists {
		return pkgErrors.Errorf("invalid function token: %s, posistion:%d",
			item.Literal, item.Position)
	}

	if funcArgs > 0 && item.FuncArgCount != funcArgs {
		return pkgErrors.Errorf(
			"function '%s' has %d arguments, expected %d, token position:%d",
			item.Literal, item.FuncArgCount, funcArgs, item.Position)
	}

	return nil
}
