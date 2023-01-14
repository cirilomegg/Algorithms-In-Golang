package stack_min

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackMin(t *testing.T) {
	stack := NewStack()

	stack.push(5)
	min, _ := stack.min()
	assert.Equal(t, 5, min)

	stack.push(3)
	min, _ = stack.min()
	assert.Equal(t, 3, min)

	stack.push(7)
	min, _ = stack.min()
	assert.Equal(t, 3, min)

	stack.push(9)
	min, _ = stack.min()
	assert.Equal(t, 3, min)

	stack.push(2)
	min, _ = stack.min()
	assert.Equal(t, 2, min)

	_, err := stack.pop()
	assert.Nil(t, err)
	min, _ = stack.min()
	assert.Equal(t, 3, min)

	_, err = stack.pop()
	assert.Nil(t, err)
	min, _ = stack.min()
	assert.Equal(t, 3, min)

	_, err = stack.pop()
	assert.Nil(t, err)
	min, _ = stack.min()
	assert.Equal(t, 3, min)

	_, err = stack.pop()
	assert.Nil(t, err)
	min, _ = stack.min()
	assert.Equal(t, 5, min)

	_, err = stack.pop()
	assert.Nil(t, err)
	_, err = stack.min()
	assert.NotNil(t, err)
}
