package decexpr

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	tests := []struct {
		exp    string
		idents map[string]decimal.Decimal
		result string
	}{
		{exp: "0", idents: map[string]decimal.Decimal{}, result: "0"},
		{exp: "-5+10", idents: map[string]decimal.Decimal{}, result: "5"},
		{exp: "sum(1, 3+5, min(5*10, 7-5))", idents: map[string]decimal.Decimal{}, result: "11"},
		{exp: "1/3", idents: map[string]decimal.Decimal{}, result: "0.3333333333333333"},
		{exp: "2^3", idents: map[string]decimal.Decimal{}, result: "8"},
		{exp: "5+3*6-val1", idents: map[string]decimal.Decimal{"val1": decimal.NewFromInt(5)}, result: "18"},
		{exp: "(0.6+0.4)+0.25*4", idents: map[string]decimal.Decimal{}, result: "2"},
		{exp: "5+-5", idents: map[string]decimal.Decimal{}, result: "0"},
		{exp: "5*5", idents: map[string]decimal.Decimal{}, result: "25"},
		{exp: "5^2", idents: map[string]decimal.Decimal{}, result: "25"},
		{exp: "5-5", idents: map[string]decimal.Decimal{}, result: "0"},
		{exp: "5/2", idents: map[string]decimal.Decimal{}, result: "2.5"},
		{exp: "5%2", idents: map[string]decimal.Decimal{}, result: "1"},
		{exp: "5+5*5", idents: map[string]decimal.Decimal{}, result: "30"},
		{exp: "(5+5)*5", idents: map[string]decimal.Decimal{}, result: "50"},
		{exp: "round(5.3555, 2)", idents: map[string]decimal.Decimal{}, result: "5.36"},
		{exp: "trunc(5.3555, 2)", idents: map[string]decimal.Decimal{}, result: "5.35"},
	}
	for _, test := range tests {
		t.Run(test.exp, func(t *testing.T) {
			v, err := Eval(test.exp, test.idents)
			assert.NoError(t, err)
			assert.Equal(t, test.result, v.String())
		})
	}
}

func TestDecimal(t *testing.T) {
	val := decimal.NewFromFloat(3123.5612)
	fmt.Println(val.Round(1).String())
	fmt.Println(val.Truncate(1).String())

	fmt.Println(true)
}

// BenchmarkEval-12    	  400081	      2658 ns/op	    3848 B/op	      47 allocs/op
// BenchmarkEval-12    	  479031	      2458 ns/op	    3848 B/op	      46 allocs/op
// BenchmarkEval-12    	  374209	      2917 ns/op	    5720 B/op	      32 allocs/op
// BenchmarkEval-12    	  435616	      2763 ns/op	    5208 B/op	      32 allocs/op
// BenchmarkEval-12    	  480310	      2467 ns/op	    4744 B/op	      41 allocs/op
// BenchmarkEval-12    	  413218	      2667 ns/op	    3160 B/op	      55 allocs/op
// BenchmarkEval-12    	  425035	      2550 ns/op	    3144 B/op	      54 allocs/op
// BenchmarkEval-12    	  462590	      2398 ns/op	    2280 B/op	      54 allocs/op
// BenchmarkEval-12    	  555511	      1895 ns/op	    1464 B/op	      35 allocs/op
func BenchmarkEval(b *testing.B) {
	b.StopTimer()
	values := map[string]decimal.Decimal{
		"val1": decimal.NewFromInt(10),
		"val2": decimal.NewFromInt(100),
		"val3": decimal.NewFromInt(50),
	}

	//exp := "sum(1, max(val1, 100), min(val2, 100)) * 20 "
	exp := "20 * 30.5 + sum(13, val1, val2, val3)/(sum(1, max(val1, 100), min(val2, 100)))"
	err := Default().ParseAndCache(exp)
	if err != nil {
		b.Fatal(err)
	}
	//f, err := os.Create("cpu.prof")
	//if err != nil {
	//	b.Fatal(err)
	//}
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v, err := Eval(exp, values)
		if err != nil {
			b.Fatal(err)
		}
		v = v
	}
}
