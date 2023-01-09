package intersection

import (
	"github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"
	"math"
)

func GetIntersection(list1 data_structure.LinkedList, list2 data_structure.LinkedList) *data_structure.Node {
	var shorter *data_structure.Node
	var longer *data_structure.Node

	length1, tail1 := getListInfo(list1)
	length2, tail2 := getListInfo(list2)

	if tail1 != tail2 {
		return nil
	}

	if length1 > length2 {
		longer, shorter = list1.Head, list2.Head
	} else {
		longer, shorter = list2.Head, list1.Head
	}

	longer = getKthNode(longer, int(math.Abs(float64(length2-length1))))

	for longer != shorter {
		shorter = shorter.Next
		longer = longer.Next
	}

	return longer
}

func getKthNode(node *data_structure.Node, k int) *data_structure.Node {
	if node == nil {
		return nil
	}

	for k > 0 && node != nil {
		node = node.Next
		k--
	}

	return node
}

func getListInfo(list data_structure.LinkedList) (int, *data_structure.Node) {
	count := 1
	node := list.Head

	for node.Next != nil {
		count++
		node = node.Next
	}

	return count, node
}
