package decexpr

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var parseFunctions = map[string]int{
	"sin":  2,
	"cos":  1,
	"tan":  1,
	"acos": 1,
	"asin": 1,
	"atan": 1,
	"sum":  -1,
	"min":  -1,
	"max":  -1,
}

func TestParser(t *testing.T) {
	parser := NewParser(parseFunctions)

	tests := []struct {
		exp    string
		output string
	}{
		{
			exp:    "3 + 4 * 2 / (1 - 5)^2",
			output: "3 4 2 * 1 5 - 2 ^ / +",
		},
		{
			exp:    "sin(2, 3)",
			output: "2 3 sin:2",
		},
		{
			exp:    "-5 * 10 / -7",
			output: "5 -. 10 * 7 -. /",
		},
		{
			exp:    "sum(1 +5 , max(3,10), min(5, -6))",
			output: "1 5 + 3 10 max:2 5 6 -. min:2 sum:3",
		},
		{
			exp:    "sum(min(1,2), max(2,3))",
			output: "1 2 min:2 2 3 max:2 sum:2",
		},

		{
			exp:    "1/3",
			output: "1 3 /",
		},
	}
	for _, test := range tests {
		t.Run(test.exp, func(t *testing.T) {
			items, err := parser.Parse(test.exp)
			assert.NoError(t, err)
			assert.EqualValues(t, test.output, sprintItems(items))
			fmt.Println(sprintItems(items))
		})
	}
}

func sprintItems(items []*RPNItem) string {
	var str strings.Builder
	for _, item := range items {
		switch item.Token.Type {
		case TokenFunction:
			str.WriteString(fmt.Sprintf("%s:%d ", item.Token.Literal, item.FuncArgCount))
		case TokenUnaryOperator:
			str.WriteString(fmt.Sprintf("%s. ", item.Token.Literal))
		default:
			str.WriteString(fmt.Sprintf("%s ", item.Token.Literal))

		}

	}

	return strings.TrimSpace(str.String())
}
