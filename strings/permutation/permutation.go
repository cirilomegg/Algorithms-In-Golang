package permutation

func IsPermutation(text1 string, text2 string) bool {
	if len(text1) == 0 || len(text1) != len(text2) {
		return false
	}

	characters := make(map[int32]int)

	for _, c := range text1 {
		characters[c]++
	}

	for _, c := range text2 {
		characters[c]--

		if characters[c] < 0 {
			return false
		}
	}
	
	return true
}
