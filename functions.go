package decexpr

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Function func(vals ...decimal.Decimal) (decimal.Decimal, error)

type FuncInfo struct {
	Call Function
	Args int
}

var functions = map[string]FuncInfo{
	"max":   {Call: Max, Args: -1},
	"min":   {Call: Min, Args: -1},
	"sum":   {Call: Sum, Args: -1},
	"round": {Call: Round, Args: -1},
}

func Max(vals ...decimal.Decimal) (decimal.Decimal, error) {
	if len(vals) == 0 {
		return decimal.Zero, nil
	}

	return decimal.Max(vals[0], vals[1:]...), nil
}

func Min(vals ...decimal.Decimal) (decimal.Decimal, error) {
	if len(vals) == 0 {
		return decimal.Zero, nil
	}

	return decimal.Min(vals[0], vals[1:]...), nil
}

func Sum(vals ...decimal.Decimal) (decimal.Decimal, error) {
	if len(vals) == 0 {
		return decimal.Zero, nil
	}

	return decimal.Sum(vals[0], vals[1:]...), nil
}

func Round(vals ...decimal.Decimal) (decimal.Decimal, error) {
	if len(vals) != 2 {
		return decimal.Zero, errors.New("invalid number of arguments")
	}

	return vals[0].Round(int32(vals[1].IntPart())), nil
}
