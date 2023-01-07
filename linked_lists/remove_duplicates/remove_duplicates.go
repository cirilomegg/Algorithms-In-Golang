package remove_duplicates

import "github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"

func RemoveDuplicates(list data_structure.LinkedList) data_structure.LinkedList {
	if list.Head == nil {
		return list
	}

	found := make(map[int]bool)
	var previous *data_structure.Node
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
