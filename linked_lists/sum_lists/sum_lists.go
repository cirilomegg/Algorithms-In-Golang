package sum_lists

import "github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"

func SumLists(list1 data_structure.LinkedList, list2 data_structure.LinkedList) data_structure.LinkedList {
	node1 := list1.Head
	node2 := list2.Head
	var result data_structure.LinkedList
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

func SumListsRecursive(list1 data_structure.LinkedList, list2 data_structure.LinkedList) data_structure.LinkedList {
	var result data_structure.LinkedList
	result.Head = AddNext(list1.Head, list2.Head, false)
	result.Tail = result.Head
	return result
}

func AddNext(node1 *data_structure.Node, node2 *data_structure.Node, carry bool) *data_structure.Node {
	var node *data_structure.Node = &data_structure.Node{}

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
