package intersection

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
	"math"
)

func GetIntersection(list1 data_structures.LinkedList, list2 data_structures.LinkedList) *data_structures.Node {
	var shorter *data_structures.Node
	var longer *data_structures.Node

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

func getKthNode(node *data_structures.Node, k int) *data_structures.Node {
	if node == nil {
		return nil
	}

	for k > 0 && node != nil {
		node = node.Next
		k--
	}

	return node
}

func getListInfo(list data_structures.LinkedList) (int, *data_structures.Node) {
	count := 1
	node := list.Head

	for node.Next != nil {
		count++
		node = node.Next
	}

	return count, node
}
