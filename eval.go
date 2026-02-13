package decexpr

import (
	pkgErrors "github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var defaultEval = NewExpressionEvaluator(true, functions)

type ExpressionEvaluator struct {
	functions map[string]FuncInfo
	parser    *Parser
	cache     EvalCache
}

func NewExpressionEvaluator(useCache bool, functions map[string]FuncInfo) *ExpressionEvaluator {
	funcArgs := make(map[string]int)
	for k, v := range functions {
		funcArgs[k] = v.Args
	}

	var evalCache EvalCache

	if useCache {
		evalCache = NewEvalMapCache()
	} else {
		evalCache = NewEvalNoopCache()
	}

	return &ExpressionEvaluator{
		functions: functions,
		parser:    NewParser(funcArgs),
		cache:     evalCache,
	}
}

func (e *ExpressionEvaluator) ParseAndCache(exp string) error {
	_, ok := e.cache.Get(exp)
	if !ok {
		items, err := e.parser.Parse(exp)
		if err != nil {
			return err
		}

		e.cache.Put(exp, items)
	}

	return nil
}

func (e *ExpressionEvaluator) Eval(exp string, identValue map[string]decimal.Decimal) (res decimal.Decimal, err error) {
	items, ok := e.cache.Get(exp)
	if !ok {
		items, err = e.parser.Parse(exp)
		if err != nil {
			return decimal.Decimal{}, err
		}

		e.cache.Put(exp, items)
	}

	res, err = e.evalRPN(items, identValue)
	if err != nil {
		return decimal.Decimal{}, pkgErrors.Wrapf(err, "invalid expression: %s", exp)
	}

	return res, nil
}

func (e *ExpressionEvaluator) evalRPN(items []*RPNItem, identValue map[string]decimal.Decimal) (decimal.Decimal, error) {
	stack := NewNumberStack(len(items))

	for len(items) > 0 {
		item := items[0]
		switch item.Type {
		case TokenNumber:
			stack.Push(item.Number)
		case TokenIdent:
			value, ok := identValue[item.Literal]
			if !ok {
				return decimal.Decimal{}, pkgErrors.Errorf(
					"ident value not found for %s, token position:%d",
					item.Literal, item.Position)
			}

			stack.Push(value)
		case TokenUnaryOperator:
			if stack.Len() < 1 {
				return decimal.Decimal{}, pkgErrors.Errorf(
					"invalid unary operator '%s', token position:%d",
					item.Literal, item.Position)
			}

			switch item.Literal {
			case "-":
				value := stack.Pop()
				stack.Push(value.Neg())
			default:
				return decimal.Decimal{}, pkgErrors.Errorf(
					"unsupported unary operator '%s', token position:%d",
					item.Literal, item.Position)
			}
		case TokenOperator:
			if stack.Len() < 2 {
				return decimal.Decimal{}, pkgErrors.Errorf(
					"invalid operator %s, token position:%d",
					item.Literal, item.Position)
			}

			var (
				val  decimal.Decimal
				val2 = stack.Pop()
				val1 = stack.Pop()
			)

			switch item.Literal {
			case "+":
				val = val1.Add(val2)
			case "-":
				val = val1.Sub(val2)
			case "*":
				val = val1.Mul(val2)
			case "/":
				if val2.IsZero() {
					return decimal.Decimal{}, pkgErrors.Errorf("division by 0, token position:%d", item.Position)
				}

				val = val1.Div(val2)
			case "%":
				if val2.IsZero() {
					return decimal.Decimal{}, pkgErrors.Errorf("division by 0, token position:%d", item.Position)
				}

				val = val1.Mod(val2)
			case "^":
				val = val1.Pow(val2)
			default:
				return decimal.Decimal{}, pkgErrors.Errorf("unsupported operator '%s', token position:%d",
					item.Literal, item.Position)
			}

			stack.Push(val)
		case TokenFunction:
			function, ok := e.functions[item.Literal]
			if !ok {
				return decimal.Decimal{}, pkgErrors.Errorf("unknown function '%s', token position:%d",
					item.Literal, item.Position)
			}

			if stack.Len() < item.FuncArgCount {
				return decimal.Decimal{}, pkgErrors.Errorf("invalid function %s, token position:%d",
					item.Literal, item.Position,
				)
			}

			vals, err := stack.PopN(item.FuncArgCount)
			if err != nil {
				return decimal.Decimal{}, err
			}

			v, err := function.Call(vals...)
			if err != nil {
				return decimal.Decimal{}, pkgErrors.Wrapf(err, "invalid function '%s', token position:%d",
					item.Literal, item.Position)
			}

			stack.Push(v)
		default:
			return decimal.Decimal{}, pkgErrors.Errorf("unknown token '%s', token position:%d",
				item.Literal, item.Position)
		}

		items = items[1:]
	}

	val := stack.Pop()

	if stack.Len() > 0 {
		return decimal.Decimal{}, pkgErrors.New("stack values is not empty")
	}

	return val, nil
}

func Eval(exp string, identValue map[string]decimal.Decimal) (decimal.Decimal, error) {
	return defaultEval.Eval(exp, identValue)
}
