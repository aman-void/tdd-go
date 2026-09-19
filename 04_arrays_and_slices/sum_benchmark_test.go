package arraysandslices

import "testing"

func BenchmarkSum(b *testing.B) {

	numbers := []int{1, 2, 3, 4, 5}
	for b.Loop() {
		Sum(numbers)

	}
}
