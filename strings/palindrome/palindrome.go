package palindrome

import "strings"

func IsPalindrome(text string) bool {
	if len(strings.TrimSpace(text)) == 0 {
		return false
	}

	text = strings.ToLower(text)

	chars := make(map[int32]int)
	oddCount := 0
	firstCharIndex := 'a'
	lastCharIndex := 'z'

	for _, charIndex := range text {
		if charIndex >= firstCharIndex && charIndex <= lastCharIndex {
			chars[charIndex]++

			if chars[charIndex]%2 != 0 {
				oddCount++
			} else {
				oddCount--
			}
		}
	}

	return oddCount <= 1
}
