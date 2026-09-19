package arraysandslices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	// t.Run("collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d want %d given %v", want, got, numbers)
	// 	}
	// })

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", want, got, numbers)
		}
	})

}

// Example
func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	fmt.Println(result)
	// Output: 15

}
