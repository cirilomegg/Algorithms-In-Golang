package unique_characters

func IsUnique(text string) bool {
	if len(text) == 0 {
		return false
	} else if len(text) == 1 {
		return true
	}

	characters := make(map[int32]bool)

	for _, c := range text {
		if characters[c] == true {
			return false
		}

		characters[c] = true
	}

	return true
}
