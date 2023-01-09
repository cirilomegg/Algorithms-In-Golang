package sum_lists

import (
	"github.com/cirilomegg/algorithms-in-golang/linked_lists/data_structure"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSumLists(t *testing.T) {
	list1, list2, expected := buildLists(0, 0)
	result := SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(1, 0)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(5, 5)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(10, 9)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(50, 50)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(612, 865)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)
}

func TestSumListsRecursive(t *testing.T) {
	list1, list2, expected := buildLists(0, 0)
	result := SumListsRecursive(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(1, 0)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(5, 5)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(10, 9)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(50, 50)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)

	list1, list2, expected = buildLists(612, 865)
	result = SumLists(list1, list2)
	assertResult(t, result, expected)
}

func assertResult(t *testing.T, result data_structure.LinkedList, expected data_structure.LinkedList) {
	node := result.Head
	expectedNode := expected.Head

	for expectedNode != nil {
		assert.NotNil(t, node)
		assert.Equal(t, expectedNode.Data, node.Data)

		expectedNode = expectedNode.Next
		node = node.Next
	}
}

func buildLists(num1 int, num2 int) (data_structure.LinkedList, data_structure.LinkedList, data_structure.LinkedList) {
	total := num1 + num2

	reversedNum1 := reverseString(strconv.Itoa(num1))
	reversedNum2 := reverseString(strconv.Itoa(num2))
	reversedTotal := reverseString(strconv.Itoa(total))

	list1 := buildList(reversedNum1)
	list2 := buildList(reversedNum2)
	expectedList := buildList(reversedTotal)

	return list1, list2, expectedList
}

func buildList(num string) data_structure.LinkedList {
	var list data_structure.LinkedList

	for _, i := range num {
		value, _ := strconv.Atoi(string(i))
		list.Append(value)
	}

	return list
}

func reverseString(text string) string {
	runes := []rune(text)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
