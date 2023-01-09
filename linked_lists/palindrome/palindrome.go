package palindrome

import "github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"

func IsPalindrome(list data_structure.LinkedList) bool {
	if list.Head == nil {
		return false
	}

	stack := data_structure.NewStack()
	node := list.Head
	runner := list.Head

	for node != nil {
		if runner != nil && runner.Next != nil {
			stack.Push(node.Data)
		} else {
			peek, _ := stack.Peek()

			if peek == node.Data {
				stack.Pop()
			}
		}

		node = node.Next

		if node != nil && node.Next != nil {
			runner = node.Next.Next
		}
	}

	return stack.IsEmpty()
}
