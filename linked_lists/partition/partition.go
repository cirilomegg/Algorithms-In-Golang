package partition

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
)

func Partition(list data_structures.LinkedList, x int) *data_structures.Node {
	node := list.Head
	head := node
	tail := node

	for node != nil {
		next := node.Next

		if node.Data < x {
			node.Next = head
			head = node
		} else {
			tail.Next = node
			tail = node
		}

		node = next
	}

	tail.Next = nil

	return head
}
