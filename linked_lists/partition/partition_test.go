package partition

import (
	"fmt"
	"github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"
	"math/rand"
	"testing"
)

func TestPartition(t *testing.T) {
	var list data_structure.LinkedList
	max := 20
	min := 0

	for i := 0; i < 20; i++ {
		list.Append(rand.Intn(max-min) + min)
	}

	head := Partition(list, 10)

	current := head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
