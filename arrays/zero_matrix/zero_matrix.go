package zero_matrix

func ZeroMatrix(matrix [][]int) [][]int {
	duplicate := duplicateMatrix(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				setZeros(duplicate, i, j)
			}
		}
	}

	return duplicate
}

func setZeros(matrix [][]int, row int, column int) {
	for j := 0; j < len(matrix[row]); j++ {
		matrix[row][j] = 0
	}

	for i := 0; i < len(matrix); i++ {
		matrix[i][column] = 0
	}
}
