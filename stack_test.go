package decexpr

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNumberStack(t *testing.T) {
	stack := NewNumberStack(10)

	val1 := decimal.NewFromInt(10)
	val2 := decimal.NewFromInt(30)
	val3 := decimal.NewFromInt(40)
	val4 := decimal.NewFromInt(60)

	stack.Push(val1)
	stack.Push(val2)
	stack.Push(val3)
	stack.Push(val4)

	vals, err := stack.PopN(5)
	assert.Error(t, err)
	assert.Nil(t, vals)

	vals, err = stack.PopN(3)
	assert.NoError(t, err)
	assert.EqualValues(t, []decimal.Decimal{val4, val3, val2}, vals)
	assert.EqualValues(t, []decimal.Decimal{val1}, stack.numbers)

	vals, err = stack.PopN(1)
	assert.NoError(t, err)
	assert.EqualValues(t, []decimal.Decimal{val1}, vals)
	assert.EqualValues(t, []decimal.Decimal{}, stack.numbers)

}

func TestSlicePointer(t *testing.T) {
	sl := []int{1, 3, 4, 5}

	v := &sl[2]
	*v++

	fmt.Println(sl)

}
