package kth_to_last_recursive

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
	kth := FindKthToLast(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	k = 10
	kth = FindKthToLast(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	k = 50
	kth = FindKthToLast(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	kth = FindKthToLast(list, 0)
	assert.Nil(t, kth)

	kth = FindKthToLast(list, 100)
	assert.Nil(t, kth)
}
