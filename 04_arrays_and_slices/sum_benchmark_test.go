package arraysandslices

import "testing"

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for b.Loop() {
		Sum(numbers)
	}
}

func BenchmarkSumLarge(b *testing.B) {
	numbers := make([]int, 10000)
	for i := range numbers {
		numbers[i] = i + 1
	}
	b.ResetTimer()
	for b.Loop() {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	c := []int{3, 4}
	for b.Loop() {
		SumAll(a, c)
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	c := []int{3, 4}
	for b.Loop() {
		SumAllTails(a, c)
	}
}
