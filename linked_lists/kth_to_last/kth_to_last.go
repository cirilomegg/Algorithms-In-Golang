package kth_to_last

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
)

func KthToLast(list data_structures.LinkedList, k int) *data_structures.Node {
	if list.Head == nil || k <= 0 {
		return nil
	}

	runner := list.Head

	for i := 0; i < k; i++ {
		if runner == nil {
			return nil
		}

		runner = runner.Next
	}

	kth := list.Head

	for runner != nil {
		kth = kth.Next
		runner = runner.Next
	}

	return kth
}

func FindKthToLastRecursive(list data_structures.LinkedList, k int) *data_structures.Node {
	kth, _ := KthToLastRecursive(list.Head, k, 0)
	return kth
}

func KthToLastRecursive(node *data_structures.Node, k int, i int) (*data_structures.Node, int) {
	if node == nil {
		return nil, i
	}

	kth, counter := KthToLastRecursive(node.Next, k, i)
	counter++

	if counter == k {
		return node, counter
	}

	return kth, counter
}
