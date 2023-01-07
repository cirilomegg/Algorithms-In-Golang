package kth_to_last_recursive

import "github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"

func FindKthToLast(list data_structure.LinkedList, k int) *data_structure.Node {
	kth, _ := KthToLast(list.Head, k, 0)
	return kth
}

func KthToLast(node *data_structure.Node, k int, i int) (*data_structure.Node, int) {
	if node == nil {
		return nil, i
	}

	kth, counter := KthToLast(node.Next, k, i)
	counter++

	if counter == k {
		return node, counter
	}

	return kth, counter
}
