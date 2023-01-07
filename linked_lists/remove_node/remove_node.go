package remove_node

import (
	"github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"
)

func RemoveNode(node *data_structure.Node) {
	if node == nil {
		return
	}

	if node.Next != nil {
		node.Data = node.Next.Data
		node.Next = node.Next.Next
	}
}
