package kth_to_last

import (
	"github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthToLast(t *testing.T) {
	var list data_structure.LinkedList
	listLength := 50

	for i := 1; i <= listLength; i++ {
		list.Append(i)
	}

	k := 1
	kth := KthToLast(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	k = 10
	kth = KthToLast(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	k = 50
	kth = KthToLast(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	kth = KthToLast(list, 0)
	assert.Nil(t, kth)

	kth = KthToLast(list, 100)
	assert.Nil(t, kth)
}
