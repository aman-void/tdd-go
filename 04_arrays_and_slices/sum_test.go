package arraysandslices

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{name: "collection of 5 numbers", numbers: []int{1, 2, 3, 4, 5}, want: 15},
		{name: "collection of any size", numbers: []int{1, 2, 3}, want: 6},
		{name: "empty slice", numbers: []int{}, want: 0},
		{name: "nil slice", numbers: nil, want: 0},
		{name: "single element", numbers: []int{7}, want: 7},
		{name: "negatives cancel out", numbers: []int{-1, -2, 3}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.numbers)
			if got != tt.want {
				t.Errorf("got %d want %d given %v", got, tt.want, tt.numbers)
			}
		})
	}
}

// Another test
func TestSumAll(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{name: "two slices", input: [][]int{{1, 2}, {0, 9}}, want: []int{3, 9}},
		{name: "three slices of different sizes", input: [][]int{{1, 2, 3}, {4}, {}}, want: []int{6, 4, 0}},
		{name: "no slices", input: nil, want: nil},
		{name: "nil and empty slices", input: [][]int{nil, {}}, want: []int{0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumAll(tt.input...)
			checkSums(t, got, tt.want)
		})
	}
}

func TestSumAllTails(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{name: "make the sums of tail of", input: [][]int{{1, 2}, {3, 4}}, want: []int{2, 4}},
		{name: "safely sum empty slices", input: [][]int{{}, {3, 4, 5}}, want: []int{0, 9}},
		{name: "single element tails are empty", input: [][]int{{1}, {3, 4, 5}}, want: []int{0, 9}},
		{name: "nil slice is safe", input: [][]int{nil, {1, 2, 3}}, want: []int{0, 5}},
		{name: "no slices", input: nil, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumAllTails(tt.input...)
			checkSums(t, got, tt.want)
		})
	}
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Example
func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	fmt.Println(result)
	// Output: 15

}

func ExampleSum_empty() {
	fmt.Println(Sum([]int{}))
	fmt.Println(Sum(nil))
	// Output:
	// 0
	// 0
}

func ExampleSumAll() {
	got := SumAll([]int{1, 2, 3, 4, 5}, []int{3, 4})
	fmt.Println(got)
	// Output: [15 7]
}

func ExampleSumAllTails() {
	got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
	fmt.Println(got)
	// Output: [5 9]
}
