package benchmarks

import (
	"exercise/slices/slices"
	"testing"
)

func BenchmarkRemoveLoop(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		slices.RemoveByIndexLoop(a, index)
	}
}

func BenchmarkRemoveNewSliceByCopies(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		slices.RemoveNewSliceByCopies(a, index)
	}
}

func BenchmarkRemoveNewSliceByCopy(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		slices.RemoveNewSliceByCopy(a, index)
	}
}

func BenchmarkRemoveByAppend(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		slices.RemoveByAppend(a, index)
	}
}

func BenchmarkRemoveBySlices(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	index := 2
	for i := 0; i < b.N; i++ {
		slices.RemoveBySlices(a, index)
	}
}
