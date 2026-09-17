package adder

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(3, 4)
	expected := 7

	if expected != sum {
		t.Errorf("expected %d but got %d", expected, sum)
	}

}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 7
}
