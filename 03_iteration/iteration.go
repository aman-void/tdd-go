package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	// A simple approach would be to concatenate strings using +=:
	//
	// var repeated string
	// for i := 0; i < repeatCount; i++ {
	//     repeated += character
	// }
	//
	// However, strings in Go are immutable. Repeatedly using += in a loop
	// can cause multiple allocations and copies as the string grows.
	//
	// For a small number of iterations this is usually insignificant,
	// but for larger loops it can become unnecessarily expensive.

	// strings.Builder is designed for efficiently constructing strings:
	//
	// var repeated strings.Builder
	// for i := 0; i < repeatCount; i++ {
	//     repeated.WriteString(character)
	// }
	//
	// return repeated.String()
	//
	// The Builder maintains an internal buffer, avoiding the repeated
	// allocation and copying that can happen with string concatenation.

	// Since our actual operation is simply "repeat this string N times",
	// the standard library already provides the most idiomatic solution.
	//
	// strings.Repeat handles the allocation and construction internally,
	// so there is no need to manually manage a Builder or a loop here.
	return strings.Repeat(character, repeatCount)
}
