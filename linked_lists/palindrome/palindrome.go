package palindrome

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
)

func IsPalindrome(list data_structures.LinkedList) bool {
	if list.Head == nil {
		return true
	}

	stack := data_structures.NewStack()
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

func IsPalindromeRecursive(list data_structures.LinkedList) bool {
	length := getLength(list)
	_, result := CheckIsPalindromeRecursive(list.Head, length)
	return result
}

func CheckIsPalindromeRecursive(node *data_structures.Node, length int) (*data_structures.Node, bool) {
	if node == nil || length <= 0 {
		return node, true
	} else if length == 1 {
		return node.Next, true
	}

	current, result := CheckIsPalindromeRecursive(node.Next, length-2)

	if !result || current == nil {
		return current, result
	}

	result = node.Data == current.Data
	current = current.Next

	return current, result
}

func getLength(list data_structures.LinkedList) int {
	count := 0
	node := list.Head

	for node != nil {
		count++
		node = node.Next
	}

	return count
}
