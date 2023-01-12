package loop_detection

import (
	"github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectLoop(t *testing.T) {
	var list data_structure.LinkedList

	for i := 0; i < 5; i++ {
		list.Append(i)
	}

	result := DetectLoop(list)
	assert.Nil(t, result)

	list.Head = nil
	list.Tail = nil
	var loopStart *data_structure.Node

	for i := 0; i < 5; i++ {
		list.Append(i)

		if i == 3 {
			loopStart = list.Tail
		}
	}

	list.Tail.Next = loopStart
	list.Tail = loopStart

	result = DetectLoop(list)
	assert.NotNil(t, result)
	assert.Equal(t, 3, result.Data)

	list.Head = nil
	list.Tail = nil

	for i := 0; i < 5; i++ {
		list.Append(i)
	}

	list.Tail.Next = list.Head
	list.Tail = list.Head

	result = DetectLoop(list)
	assert.NotNil(t, result)
	assert.Equal(t, 0, result.Data)
}
