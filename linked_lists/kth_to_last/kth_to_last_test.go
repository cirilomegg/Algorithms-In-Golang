package kth_to_last

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthToLast(t *testing.T) {
	var list data_structures.LinkedList
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

func TestKthToLastRecursive(t *testing.T) {
	var list data_structures.LinkedList
	listLength := 50

	for i := 1; i <= listLength; i++ {
		list.Append(i)
	}

	k := 1
	kth := FindKthToLastRecursive(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	k = 10
	kth = FindKthToLastRecursive(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	k = 50
	kth = FindKthToLastRecursive(list, k)
	assert.Equal(t, listLength-k+1, kth.Data)

	kth = FindKthToLastRecursive(list, 0)
	assert.Nil(t, kth)

	kth = FindKthToLastRecursive(list, 100)
	assert.Nil(t, kth)
}
