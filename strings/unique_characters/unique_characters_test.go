package unique_characters

import "testing"

func TestIsUnique(t *testing.T) {
	data := getSampleData()

	for key, value := range data {
		isUnique := IsUnique(key)

		if isUnique != value {
			t.Errorf("The IsUnique should return %v for text '%v', Got: %v, expected: %v", value, key, isUnique, value)
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
