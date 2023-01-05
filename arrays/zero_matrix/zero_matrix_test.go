package zero_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZero(t *testing.T) {
	var matrix = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	var expectedMatrix = [][]int{
		{0, 0, 0},
		{0, 4, 5},
		{0, 7, 8},
	}
	evaluate(t, matrix, expectedMatrix)

	matrix = [][]int{
		{0, 0, 0},
		{3, 4, 5},
		{6, 7, 8},
	}
	expectedMatrix = [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	evaluate(t, matrix, expectedMatrix)

	matrix = [][]int{
		{1, 0, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	expectedMatrix = [][]int{
		{0, 0, 0},
		{3, 0, 5},
		{6, 0, 8},
	}
	evaluate(t, matrix, expectedMatrix)

	matrix = [][]int{
		{1, 5, 2},
		{3, 0, 5},
		{6, 7, 8},
	}
	expectedMatrix = [][]int{
		{1, 0, 2},
		{0, 0, 0},
		{6, 0, 8},
	}
	evaluate(t, matrix, expectedMatrix)

	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 9},
	}
	expectedMatrix = [][]int{
		{1, 0, 3},
		{4, 0, 6},
		{0, 0, 0},
	}
	evaluate(t, matrix, expectedMatrix)

	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	expectedMatrix = [][]int{
		{1, 2, 0},
		{4, 5, 0},
		{0, 0, 0},
	}
	evaluate(t, matrix, expectedMatrix)
}

func evaluate(t *testing.T, matrix [][]int, expectedMatrix [][]int) {
	matrix = ZeroMatrix(matrix)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			assert.Equal(t, expectedMatrix[row][col], matrix[row][col])
		}
	}
}
