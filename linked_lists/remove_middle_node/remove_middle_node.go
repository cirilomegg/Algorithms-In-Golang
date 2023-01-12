package remove_middle_node

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
)

func RemoveMiddleNode(list data_structures.LinkedList) data_structures.LinkedList {
	if list.Head == nil || list.Head.Next == nil {
		return list
	}

	var runner *data_structures.Node
	var middle *data_structures.Node
	var previous *data_structures.Node

	middle = list.Head
	runner = list.Head

	for runner != list.Tail {
		previous = middle
		middle = middle.Next

		if runner.Next != nil {
			if runner.Next.Next != nil {
				runner = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
	}

	if middle != list.Head && middle != list.Tail {
		previous.Next = middle.Next
	}

	return list
}
