package decexpr

import (
	pkgErrors "github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const (
	defaultInnerFunction = 4
)

type StackItems struct {
	Items []*RPNItem
}

func NewStackItems() *StackItems {
	return &StackItems{
		Items: make([]*RPNItem, 0, 10),
	}
}

func (s *StackItems) Push(item *RPNItem) {
	s.Items = append(s.Items, item)
}

func (s *StackItems) Pop() *RPNItem {
	if len(s.Items) == 0 {
		return nil
	}

	item := s.Items[len(s.Items)-1]

	s.Items = s.Items[:len(s.Items)-1]

	return item
}

func (s *StackItems) Peek() *RPNItem {
	if len(s.Items) == 0 {
		return nil
	}

	return s.Items[len(s.Items)-1]
}

func (s *StackItems) Len() int {
	return len(s.Items)
}

type ArgStack struct {
	args []int
}

func NewArgStack() *ArgStack {
	return &ArgStack{
		args: make([]int, 0, defaultInnerFunction),
	}
}

func (s *ArgStack) Push() {
	s.args = append(s.args, 1)
}

func (s *ArgStack) Pop() int {
	if len(s.args) == 0 {
		return 0
	}

	v := s.args[len(s.args)-1]

	s.args = s.args[:len(s.args)-1]

	return v
}

func (s *ArgStack) Inc() {
	if len(s.args) == 0 {
		return
	}

	s.args[len(s.args)-1]++
}

func (s *ArgStack) Len() int {
	return len(s.args)
}

type NumberStack struct {
	numbers []decimal.Decimal
}

func NewNumberStack(size int) *NumberStack {
	return &NumberStack{
		numbers: make([]decimal.Decimal, 0, size),
	}
}

func (s *NumberStack) Push(value decimal.Decimal) {
	s.numbers = append(s.numbers, value)
}

func (s *NumberStack) Pop() decimal.Decimal {
	if len(s.numbers) == 0 {
		return decimal.Decimal{}
	}

	v := s.numbers[len(s.numbers)-1]

	s.numbers = s.numbers[:len(s.numbers)-1]

	return v
}

func (s *NumberStack) PopN(n int) ([]decimal.Decimal, error) {
	if len(s.numbers) < n {
		return nil, pkgErrors.New("Number stack is too short")
	}

	res := make([]decimal.Decimal, 0, n)
	for i := len(s.numbers) - 1; i >= len(s.numbers)-n; i-- {
		res = append(res, s.numbers[i])
	}

	s.numbers = s.numbers[:len(s.numbers)-n]

	return res, nil
}

func (s *NumberStack) Len() int {
	return len(s.numbers)
}
