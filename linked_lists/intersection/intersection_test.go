package intersection

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIntersection(t *testing.T) {
	list1 := buildList([]int{1, 2, 3, 4, 5, 6, 7})
	list2 := buildList([]int{3, 2, 1})

	node := getKthNode(list1.Head, 5)
	list2.Tail.Next = node
	list2.Tail = node

	result := GetIntersection(list1, list2)
	assert.NotNil(t, result)
	assert.Equal(t, 6, result.Data)

	list1 = buildList([]int{1, 2, 3, 4, 5, 6, 7})
	list2 = buildList([]int{3, 2, 1})

	result = GetIntersection(list1, list2)
	assert.Nil(t, result)
}

func buildList(values []int) data_structures.LinkedList {
	var list data_structures.LinkedList

	for _, value := range values {
		list.Append(value)
	}

	return list
}
