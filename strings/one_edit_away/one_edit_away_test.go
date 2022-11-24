package one_edit_away

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	Text1    string
	Text2    string
	Expected bool
}

func TestIsOneEditAway(t *testing.T) {
	data := [6]TestData{
		{Text1: "abc", Text2: "abcdefgh", Expected: false},
		{Text1: "abc", Text2: "acd", Expected: false},
		{Text1: "cake", Text2: "bake", Expected: true},
		{Text1: "bake", Text2: "baked", Expected: true},
		{Text1: "baked", Text2: "bake", Expected: true},
		{Text1: "cake", Text2: "cake", Expected: false},
	}

	for _, item := range data {
		isOneEditAway := IsOneEditAway(item.Text1, item.Text2)
		assert.Equal(t, isOneEditAway, item.Expected, "Items should be equal '%v', '%v'", item.Text1, item.Text2)
	}
}
