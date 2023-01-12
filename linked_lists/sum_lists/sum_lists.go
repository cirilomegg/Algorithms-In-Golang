package sum_lists

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
)

func SumLists(list1 data_structures.LinkedList, list2 data_structures.LinkedList) data_structures.LinkedList {
	node1 := list1.Head
	node2 := list2.Head
	var result data_structures.LinkedList
	carry := false

	for node1 != nil || node2 != nil {
		count := 0

		if node1 != nil {
			count += node1.Data
			node1 = node1.Next
		}

		if node2 != nil {
			count += node2.Data
			node2 = node2.Next
		}

		if carry {
			count++
			carry = false
		}

		result.Append(count % 10)

		if count >= 10 {
			carry = true
		}
	}

	if carry {
		result.Append(1)
	}

	return result
}

func SumListsRecursive(list1 data_structures.LinkedList, list2 data_structures.LinkedList) data_structures.LinkedList {
	var result data_structures.LinkedList
	result.Head = AddNext(list1.Head, list2.Head, false)
	result.Tail = result.Head
	return result
}

func AddNext(node1 *data_structures.Node, node2 *data_structures.Node, carry bool) *data_structures.Node {
	var node *data_structures.Node = &data_structures.Node{}

	if node1 == nil && node2 == nil {
		if carry {
			node.Data = 1
		}

		return node
	}

	count := 0

	if carry {
		count++
	}

	if node1 != nil {
		count += node1.Data
		node1 = node1.Next
	}

	if node2 != nil {
		count += node2.Data
		node2 = node2.Next
	}

	node.Data = count % 10
	carry = count > 10
	node.Next = AddNext(node1, node2, carry)

	return node
}
