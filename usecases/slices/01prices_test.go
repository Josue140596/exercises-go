package slices

import "testing"

//Benchmarks

func BenchmarkLoop(b *testing.B) {
	prices := []int{50, 30, 20, 15, 25}

	for i := 0; i < b.N; i++ {
		CalculateTotalPurchase(prices)
	}
}
