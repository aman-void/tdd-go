package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("*")
	fmt.Println(repeated)
	// Output: *****
}

// strings compare (Table Driven Test)
func TestCompareStrings(t *testing.T) {

	tests := []struct {
		name     string
		str1     string
		str2     string
		expected int
	}{
		{
			name:     "equal strings",
			str1:     "hi",
			str2:     "hi",
			expected: 0,
		},
		{
			name:     "first string is smaller",
			str1:     "apple",
			str2:     "watermelon",
			expected: -1,
		},
		{
			name:     "first string is greater",
			str1:     "california",
			str2:     "boston",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareStrings(tt.str1, tt.str2)

			if result != tt.expected {
				t.Errorf(
					"CompareStrings(%q, %q) = %d and expected %d",
					tt.str1,
					tt.str2,
					result,
					tt.expected,
				)
			}
		})
	}

}
