package arraysandslices

// If we pass `numbers [5]int` to the Sum func parameter
// it works for array but not for slice. So we are using slice here.

func Sum(numbers []int) int {

	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
