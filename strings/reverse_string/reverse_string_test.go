package reverse_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	text := ""
	reverse := ReverseString(text)
	expected := ""
	assert.Equal(t, expected, reverse)

	text = "a"
	reverse = ReverseString(text)
	expected = "a"
	assert.Equal(t, expected, reverse)

	text = "water"
	reverse = ReverseString(text)
	expected = "retaw"
	assert.Equal(t, expected, reverse)
}
