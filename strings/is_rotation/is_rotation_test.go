package is_rotation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsRotation(t *testing.T) {
	s1 := "unreasonable"
	s2 := "ableunreason"
	result := IsRotation(s1, s2)
	assert.Equal(t, true, result)

	s1 = "unreasonable"
	s2 = "ablereason"
	result = IsRotation(s1, s2)
	assert.Equal(t, false, result)

	s1 = "unreasonable"
	s2 = "unablereason"
	result = IsRotation(s1, s2)
	assert.Equal(t, false, result)
}
