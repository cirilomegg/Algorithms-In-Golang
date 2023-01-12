package palindrome

import (
	"github.com/cirilomegg/algorithms-in-golang/data_structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	list := BuildList([]int{})
	isPalindrome := IsPalindrome(list)
	assert.Equal(t, true, isPalindrome)

	list = BuildList([]int{1})
	isPalindrome = IsPalindrome(list)
	assert.Equal(t, true, isPalindrome)

	list = BuildList([]int{1, 2, 3, 2, 1})
	isPalindrome = IsPalindrome(list)
	assert.Equal(t, true, isPalindrome)

	list = BuildList([]int{1, 2, 3, 3, 2, 1})
	isPalindrome = IsPalindrome(list)
	assert.Equal(t, true, isPalindrome)

	list = BuildList([]int{1, 2, 3, 4, 2, 1})
	isPalindrome = IsPalindrome(list)
	assert.Equal(t, false, isPalindrome)
}

func TestIsPalindromeRecursive(t *testing.T) {
	list := BuildList([]int{})
	isPalindrome := IsPalindromeRecursive(list)
	assert.Equal(t, true, isPalindrome)

	list = BuildList([]int{1})
	isPalindrome = IsPalindromeRecursive(list)
	assert.Equal(t, true, isPalindrome)

	list = BuildList([]int{1, 2, 3, 2, 1})
	isPalindrome = IsPalindromeRecursive(list)
	assert.Equal(t, true, isPalindrome)

	list = BuildList([]int{1, 2, 3, 3, 2, 1})
	isPalindrome = IsPalindromeRecursive(list)
	assert.Equal(t, true, isPalindrome)

	list = BuildList([]int{1, 2, 3, 4, 2, 1})
	isPalindrome = IsPalindromeRecursive(list)
	assert.Equal(t, false, isPalindrome)
}

func BuildList(values []int) data_structures.LinkedList {
	var list data_structures.LinkedList

	for _, value := range values {
		list.Append(value)
	}

	return list
}
