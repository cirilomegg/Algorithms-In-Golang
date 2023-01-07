package remove_node

import (
	"github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNode(t *testing.T) {
	var list data_structure.LinkedList
	var removedNode *data_structure.Node
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
