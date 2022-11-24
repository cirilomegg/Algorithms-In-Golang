package permutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	Text1    string
	Text2    string
	Expected bool
}

func TestIsPermutation(t *testing.T) {
	data := getSampleData()

	for _, value := range data {
		isPermutation := IsPermutation(value.Text1, value.Text2)

		if isPermutation != value.Expected {
			assert.Equal(t, value.Expected, isPermutation)
		}
	}
}

func getSampleData() []TestData {
	return []TestData{
		{Text1: "", Text2: "", Expected: false},
		{Text1: "dog", Text2: "god", Expected: true},
		{Text1: "dog", Text2: "god ", Expected: false},
		{Text1: "abba", Text2: "baab", Expected: true},
		{Text1: "aaab", Text2: "aabb", Expected: false},
		{Text1: "ab ba", Text2: "ba ab", Expected: true},
		{Text1: "abbaccba", Text2: "baabcc", Expected: false},
	}
}
