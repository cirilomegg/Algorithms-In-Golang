package remove_duplicates

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var list data_structures.LinkedList

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			list.Append(j)
		}
	}

	list = RemoveDuplicates(list)

	expected := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	current := list.Head
	count := 0

	for current != nil {
		assert.Equal(t, expected[count], current.Data)
		current = current.Next
		count++
	}
}
