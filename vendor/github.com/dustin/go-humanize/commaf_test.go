// +build go1.6

package humanize

import (
	"math"
	"math/big"
	"testing"
)

func BenchmarkBigCommaf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Commaf(1234567890.83584)
	}
}

func TestBigCommafs(t *testing.T) {
	testList{
		{"0", BigCommaf(big.NewFloat(0)), "0"},
		{"10.11", BigCommaf(big.NewFloat(10.11)), "10.11"},
		{"100", BigCommaf(big.NewFloat(100)), "100"},
		{"1,000", BigCommaf(big.NewFloat(1000)), "1,000"},
		{"10,000", BigCommaf(big.NewFloat(10000)), "10,000"},
		{"100,000", BigCommaf(big.NewFloat(100000)), "100,000"},
		{"834,142.32", BigCommaf(big.NewFloat(834142.32)), "834,142.32"},
		{"10,000,000", BigCommaf(big.NewFloat(10000000)), "10,000,000"},
		{"10,100,000", BigCommaf(big.NewFloat(10100000)), "10,100,000"},
		{"10,010,000", BigCommaf(big.NewFloat(10010000)), "10,010,000"},
		{"10,001,000", BigCommaf(big.NewFloat(10001000)), "10,001,000"},
		{"123,456,789", BigCommaf(big.NewFloat(123456789)), "123,456,789"},
		{"maxf64", BigCommaf(big.NewFloat(math.MaxFloat64)), "179,769,313,486,231,570,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000"},
		{"minf64", BigCommaf(big.NewFloat(math.SmallestNonzeroFloat64)), "0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004940656458412465"},
		{"-123,456,789", BigCommaf(big.NewFloat(-123456789)), "-123,456,789"},
		{"-10,100,000", BigCommaf(big.NewFloat(-10100000)), "-10,100,000"},
		{"-10,010,000", BigCommaf(big.NewFloat(-10010000)), "-10,010,000"},
		{"-10,001,000", BigCommaf(big.NewFloat(-10001000)), "-10,001,000"},
		{"-10,000,000", BigCommaf(big.NewFloat(-10000000)), "-10,000,000"},
		{"-100,000", BigCommaf(big.NewFloat(-100000)), "-100,000"},
		{"-10,000", BigCommaf(big.NewFloat(-10000)), "-10,000"},
		{"-1,000", BigCommaf(big.NewFloat(-1000)), "-1,000"},
		{"-100.11", BigCommaf(big.NewFloat(-100.11)), "-100.11"},
		{"-10", BigCommaf(big.NewFloat(-10)), "-10"},
	}.validate(t)
}
