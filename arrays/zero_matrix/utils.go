package zero_matrix

func duplicateMatrix(matrix [][]int) [][]int {
	duplicate := make([][]int, len(matrix))

	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}
