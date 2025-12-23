package filter

func FilterPositive(numbers []int) []int {
	var result []int
	for _, number := range numbers {
		if number > 0 {
			result = append(result, number)
		}
	}
	return result
}
