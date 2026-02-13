package decexpr

import (
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

// BenchmarkDecimalFromString-12    	 6768950	       170.1 ns/op	      48 B/op	       3 allocs/op
func BenchmarkDecimalFromString(b *testing.B) {
	str := "123.123"

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v, err := decimal.NewFromString(str)
		if err != nil {
			b.Fatal(err)
		}
		v = v
	}
}

// BenchmarkDecimalFromFloat64-12    	 2342931	       503.0 ns/op	      40 B/op	       2 allocs/op
func BenchmarkDecimalFromFloat64(b *testing.B) {
	f := 123.123

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := decimal.NewFromFloat(f)
		v = v
	}
}

// BenchmarkDecimalFromInt-12    	1000000000	         0.2663 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDecimalFromInt(b *testing.B) {
	i64 := int64(123123)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := decimal.NewFromInt(i64)
		v = v
	}
}

// BenchmarkDecimalNew-12    	1000000000	         0.2712 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDecimalNew(b *testing.B) {
	i64 := int64(123123)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := decimal.New(i64, -3)
		v = v
	}
}

// BenchmarkAtoi-12    	174426709	         6.975 ns/op	       0 B/op	       0 allocs/op
func BenchmarkAtoi(b *testing.B) {
	str := "123123"
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		v, err := strconv.Atoi(str)
		if err != nil {
			b.Fatal(err)
		}

		v = v
	}

}
