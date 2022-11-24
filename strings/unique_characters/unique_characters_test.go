package unique_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUnique(t *testing.T) {
	data := getSampleData()

	for key, value := range data {
		isUnique := IsUnique(key)

		if isUnique != value {
			assert.Equal(t, value, isUnique)
		}
	}
}

func getSampleData() map[string]bool {
	return map[string]bool{
		"":                           false,
		" ":                          true,
		"      ":                     false,
		"abcdefghijklmnopqrstuvwxyz": true,
		"abbabaacbacd":               false,
		"ab cd":                      true,
		"a b c d":                    false,
	}
}
