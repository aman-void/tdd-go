package iseven

import (
	"fmt"
	"testing"
)

func TestIsEven(t *testing.T) {
	isEven := IsEven(10)
	expected := true

	if expected != isEven {
		t.Errorf("expected %t but got %t", expected, isEven)
	}
}

func ExampleIsEven_odd() {
	fmt.Println(IsEven(7))
	// Output: false
}

func ExampleIsEven_even() {
	fmt.Println(IsEven(8))
	// Output: true
}

func ExampleIsEven_zero() {
	fmt.Println(IsEven(0))
	// Output: true
}

func ExampleIsEven_negative() {
	fmt.Println(IsEven(-2))
	// Output: true
}
