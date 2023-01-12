package loop_detection

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
)

func DetectLoop(list data_structures.LinkedList) *data_structures.Node {
	if list.Head == nil || list.Head.Next == nil {
		return nil
	}

	slowRunner := list.Head
	fastRunner := list.Head

	for fastRunner != nil && fastRunner.Next != nil {
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next

		if slowRunner == fastRunner {
			break
		}
	}

	if fastRunner == nil || fastRunner.Next == nil {
		return nil
	}

	slowRunner = list.Head

	for slowRunner != fastRunner {
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next
	}

	return fastRunner
}
