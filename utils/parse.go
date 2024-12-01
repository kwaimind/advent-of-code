package utils

func UnzipPairs(input [][]int) ([]int, []int) {
	n := len(input)
	left := make([]int, 0, n)
	right := make([]int, 0, n)

	for _, row := range input {
		left = append(left, row[0])
		right = append(right, row[1])
	}

	return left, right
}