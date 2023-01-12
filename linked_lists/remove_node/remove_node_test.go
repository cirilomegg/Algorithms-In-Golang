package remove_node

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNode(t *testing.T) {
	var list data_structures.LinkedList
	var removedNode *data_structures.Node
	removedNodeIndex := 5

	for i := 1; i <= 10; i++ {
		list.Append(i)

		if i == removedNodeIndex {
			removedNode = list.Tail
		}
	}

	RemoveNode(removedNode)

	current := list.Head

	for current != nil {
		assert.NotEqual(t, removedNodeIndex, current.Data)
		current = current.Next
	}
}
