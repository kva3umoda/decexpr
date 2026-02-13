# decexpr
decexpr is a high-precision mathematical expression evaluator for Go that uses 
[github.com/shopspring/decimal](github.com/shopspring/decimal) under the hood to eliminate floating-point arithmetic errors. Expressions are compiled to Reverse Polish Notation (RPN) for efficient and accurate evaluation with support for variables, nested functions, and operator precedence.

# Features
✅ Arbitrary precision decimal arithmetic using shopspring/decimal
✅ Full operator precedence support (+, -, *, /, %, ^)
✅ Variables (val1, price, tax_rate, etc.)
✅ Nested function calls (sum(1, min(5, 10), 3*val))
✅ Built-in functions: sum, min, max, abs, round, ceil, floor
✅ Custom functions registration
✅ RPN compilation for fast repeated evaluations
✅ Detailed error messages with position indicators
✅ Zero dependencies beyond shopspring/decimal


# Installation

```bash
go get github.com/kva3umoda/decexpr
```

# Quick Start

```go
package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/kva3umoda/decexpr"
)

func main() {
	// Create a new evaluator
	evaluator := decexpr.New()

	// Define variables
	vars := map[string]decimal.Decimal{
		"val1": decimal.NewFromInt(10),
		"val2": decimal.NewFromInt(3),
	}

	// Evaluate expression with modulo operator
	expr := "sum(1, 3+5, min(5*10, 7-val1)) + val2 * 6.67 + (10 % 3)"
	result, err := evaluator.Evaluate(expr, vars)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %s\n", result.String()) // Result: 40.01
}
```


# Built-in Functions

|Function|Description|Example|
|---|---|---|
|sum|Sum of numbers|sum(1, 2, 3) → 6|
|min|Minimum value|min(5, 2, 8) → 2|
|max|Maximum value|max(5, 2, 8) → 8|
|abs|Absolute value|abs(-5) → 5|
|round|Round to nearest integer|round(3.7) → 4|
|floor|Round down|floor(3.7) → 3|
|ceil|Round up|ceil(3.2) → 4|


## Custom Functions

You can easily add your own functions:

```go 
evaluator := decimaleval.New()

// Add power function
evaluator.AddFunction("pow", func(vals ...decimal.Decimal) (decimal.Decimal, error) {
if len(vals) != 2 {
return decimal.Decimal{}, fmt.Errorf("pow requires 2 arguments, got %d", len(vals))
}
return vals[0].Pow(vals[1]), nil
})

// Add average function
evaluator.AddFunction("avg", func(vals ...decimal.Decimal) (decimal.Decimal, error) {
if len(vals) == 0 {
return decimal.Decimal{}, fmt.Errorf("avg requires at least 1 argument")
}
sum := decimal.Zero
for _, v := range vals {
sum = sum.Add(v)
}
return sum.Div(decimal.NewFromInt(int64(len(vals)))), nil
})

result, _ := evaluator.Evaluate("pow(2, 3) + avg(10, 20, 30)", nil)
fmt.Println(result) // 28
```

## Supported Operators

|Operator|Description| Example   |
|---|---|---|
|+|Addition|5 + 3 → 8|
|-|Subtraction|5 - 3 → 2|
|*|Multiplication|5 * 3 → 15|
|/|Division|6 / 3 → 2|
|^|Power|10 % 3 → 1|
|%|Modulo (remainder)|2 ^ 3 → 8|
|()|Parentheses for grouping|(2 + 3) * 4 → 20|


## Expression Examples

```go

evaluator := decimaleval.New()

// Simple arithmetic with modulo
result, _ := evaluator.Evaluate("(100 + 20) % 30", nil)
fmt.Println(result) // 0 (120 % 30 = 0)

// Nested functions with variables
vars := map[string]decimal.Decimal{
    "discount": decimal.NewFromFloat(0.15),
    "price":    decimal.NewFromInt(100),
}
result, _ = evaluator.Evaluate("price * (1 - discount) + tax(price * 0.1)", vars)
fmt.Println(result)

// Complex expression with modulo
result, _ = evaluator.Evaluate("sum(10 % 3, 20 % 7, min(15, 30 % 4))", nil)
fmt.Println(result) // (1 + 6 + min(15, 2)) = 1 + 6 + 2 = 9
```

# Performance
The library uses Reverse Polish Notation (RPN) for efficient expression evaluation. Expressions are parsed once and can be evaluated multiple times with different variable values.

# Error Handling
The library provides detailed error messages for:
* Syntax errors
* Undefined variables
* Invalid function calls (wrong number of arguments)
* Division by zero
* Modulo by zero
* Type mismatches
  
# License
MIT License - see [LICENSE](./LICENSE) file for details.


# Why decexpr?
Traditional floating-point arithmetic can lead to precision issues in financial and scientific calculations. By using shopspring/decimal, this library ensures accurate results every time. The RPN-based evaluation engine provides fast and reliable expression parsing and evaluation, while the modulo operator adds flexibility for various mathematical operations.