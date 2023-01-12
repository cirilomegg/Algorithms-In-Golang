package remove_middle_node

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveMiddleNode(t *testing.T) {
	list := buildList(1)
	list = RemoveMiddleNode(list)
	assert.Equal(t, 1, list.Head.Data)
	assert.Equal(t, 1, list.Tail.Data)

	list = buildList(2)
	list = RemoveMiddleNode(list)
	assert.Equal(t, 1, list.Head.Data)
	assert.Equal(t, 2, list.Tail.Data)

	listLengths := [3]int{10, 20, 50}

	for _, length := range listLengths {
		list = buildList(length)
		list = RemoveMiddleNode(list)
		assertRemoved(t, list, length/2+1)
	}
}

func buildList(length int) data_structures.LinkedList {
	var list data_structures.LinkedList

	for i := 1; i <= length; i++ {
		list.Append(i)
	}

	return list
}

func assertRemoved(t *testing.T, list data_structures.LinkedList, removedNode int) {
	current := list.Head

	for current != nil {
		assert.NotEqual(t, removedNode, current.Data)
		current = current.Next
	}
}
