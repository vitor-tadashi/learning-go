package arrays_slices

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(args ...[]int) []int {
	var result []int

	for _, arg := range args {
		result = append(result, Sum(arg))
	}

	return result
}

func SumTails(args ...[]int) []int {
	var result []int

	for _, arg := range args {
		if len(arg) == 0 {
			result = append(result, 0)
			continue
		}

		result = append(result, Sum(arg[1:]))
	}

	return result
}
