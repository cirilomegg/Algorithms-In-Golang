package partition

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	var list data_structures.LinkedList

	for i := 19; i >= 0; i-- {
		list.Append(i)
	}

	head := Partition(list, 10)

	var expected data_structures.LinkedList

	for i := 0; i < 10; i++ {
		expected.Append(i)
	}

	for i := 19; i >= 10; i-- {
		expected.Append(i)
	}

	node := head
	expectedNode := expected.Head

	for node != nil {
		assert.Equal(t, expectedNode.Data, node.Data)
		expectedNode = expectedNode.Next
		node = node.Next
	}
}
