package replace_spaces

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReplaceSpaces(t *testing.T) {
	url := "Some entry to replace      "
	result := ReplaceSpaces(url, len(strings.TrimSpace(url)))

	assert.Equal(t, "Some%20entry%20to%20replace", result)
}
