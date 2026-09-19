package arraysandslices

// If we pass `numbers [5]int` to the Sum func parameter
// it works for array but not for slice. So we are using slice here.

// Sum returns the total of a slice of ints.
// An empty or nil slice sums to 0.
func Sum(numbers []int) int {

	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// SumAll returns the sum of each slice passed as a variadic argument.
// For example, SumAll([]int{1, 2}, []int{0, 9}) returns []int{3, 9}.
func SumAll(numbersToSum ...[]int) []int {

	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	var sums []int
	for _, numbers := range numbersToSum {

		// sums[i] = Sum(numbers)
		// more efficient using append

		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails returns the sum of each slice's tail (all elements except the first).
// Empty slices and single-element slices contribute 0, so no panic on numbers[1:].
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
