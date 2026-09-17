package subtractor

import (
	"fmt"
	"testing"
)

func TestSubtract(t *testing.T) {
	sub := Subtract(12, 3)
	expected := 9

	if expected != sub {
		t.Errorf("expected %d but got %d", expected, sub)
	}
}

func ExampleSubtract() {
	sub := Subtract(12, 4)
	fmt.Println(sub)
	// Output: 8
}
