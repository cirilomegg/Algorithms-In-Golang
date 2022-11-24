package palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	data := make(map[string]bool)
	data[""] = false
	data[" "] = false
	data["dog"] = false
	data["good"] = false
	data["aba"] = true
	data["Was it a car or a cat i saw?"] = true
	data["chicken ne kc ihc"] = true
	data["chicken ne kc ih c"] = true
	data["chicken ne kc i hc"] = true
	data["chicken ne k cih c"] = true
	data["chicken ne k ci hc"] = true
	data["chicken ne k c ihc"] = true
	data["chicken n ek cih c"] = true
	data["chicken n ek ci hc"] = true
	data["chicken n ek c ihc"] = true
	data["chicken n e kc ihc"] = true
	data["chicken ek cih c"] = true
	data["chicken ek ci hc"] = true
	data["chicken ek c ihc"] = true
	data["chicken e kc ihc"] = true

	for key, value := range data {
		isPalindrome := IsPalindrome(key)

		if isPalindrome != value {
			assert.Equal(t, value, isPalindrome)
		}
	}
}
