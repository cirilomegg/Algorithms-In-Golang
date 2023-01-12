package remove_node

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
)

func RemoveNode(node *data_structures.Node) {
	if node == nil {
		return
	}

	if node.Next != nil {
		node.Data = node.Next.Data
		node.Next = node.Next.Next
	}
}
