package one_edit_away

import "math"

func IsOneEditAway(first string, second string) bool {
	if first == second || math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}

	string1 := first
	string2 := second

	if len(first) > len(second) {
		string1 = second
		string2 = first
	}

	index1 := 0
	index2 := 0
	string1Length := len(string1)
	string2Length := len(string2)
	foundDifference := false

	for index1 < string1Length && index2 < string2Length {
		if string1[index1] != string2[index2] {
			if foundDifference {
				return false
			}

			foundDifference = true

			if string1Length == string2Length {
				index1++
			}
		} else {
			index1++
		}

		index2++
	}

	return true
}
