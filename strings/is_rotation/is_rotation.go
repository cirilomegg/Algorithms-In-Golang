package is_rotation

import "strings"

func IsRotation(s1 string, s2 string) bool {
	len1 := len(s1)
	len2 := len(s2)

	if len1 == len2 && len1 > 0 {
		s1s1 := s1 + s1
		return strings.Contains(s1s1, s2)
	}

	return false
}
