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
	"avg":   {Call: Avg, Args: -1},
	"round": {Call: Round, Args: 2},
	"floor": {Call: Floor, Args: 1},
	"ceil":  {Call: Ceil, Args: 1},
	"abs":   {Call: Abs, Args: 1},
	"trunc": {Call: Trunc, Args: 2},
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
	switch len(vals) {
	case 1:
		return vals[0].Round(0), nil
	case 2:
		if vals[1].Exponent() < 0 {
			return decimal.Zero, errors.New("invalid second of arguments")
		}

		return vals[0].Round(int32(vals[1].IntPart())), nil
	default:
		return decimal.Zero, errors.New("invalid number of arguments")
	}
}

func Abs(vals ...decimal.Decimal) (decimal.Decimal, error) {
	if len(vals) != 1 {
		return decimal.Zero, errors.New("invalid number of arguments")
	}

	return vals[0].Abs(), nil
}

func Floor(vals ...decimal.Decimal) (decimal.Decimal, error) {
	if len(vals) != 1 {
		return decimal.Zero, errors.New("invalid number of arguments")
	}

	return vals[0].Floor(), nil
}

func Ceil(vals ...decimal.Decimal) (decimal.Decimal, error) {
	if len(vals) != 1 {
		return decimal.Zero, errors.New("invalid number of arguments")
	}
	return vals[0].Ceil(), nil
}

func Trunc(vals ...decimal.Decimal) (decimal.Decimal, error) {
	switch len(vals) {
	case 1:
		return vals[0].Truncate(0), nil
	case 2:
		if vals[1].Exponent() < 0 {
			return decimal.Zero, errors.New("invalid second of arguments")
		}

		return vals[0].Truncate(int32(vals[1].IntPart())), nil
	default:
		return decimal.Zero, errors.New("invalid number of arguments")
	}
}

func Avg(vals ...decimal.Decimal) (decimal.Decimal, error) {
	if len(vals) == 0 {
		return decimal.Zero, nil
	}

	return decimal.Avg(vals[0], vals[1:]...), nil
}
