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
