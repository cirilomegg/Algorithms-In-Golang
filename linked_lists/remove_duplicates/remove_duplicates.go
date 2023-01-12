package remove_duplicates

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
)

func RemoveDuplicates(list data_structures.LinkedList) data_structures.LinkedList {
	if list.Head == nil {
		return list
	}

	found := make(map[int]bool)
	var previous *data_structures.Node
	current := list.Head

	for current != nil {
		if found[current.Data] == true {
			previous.Next = current.Next
		} else {
			previous = current
		}

		found[current.Data] = true
		current = current.Next
	}

	return list
}
