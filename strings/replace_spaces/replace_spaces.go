package replace_spaces

func ReplaceSpaces(text string, length int) string {
	outputLength := length + (getSpacesCount(text, 0, length) * 2)
	outputText := make([]uint8, outputLength)

	outputIndex := 0

	for i := 0; i < length; i++ {
		if text[i] == ' ' {
			outputText[outputIndex] = '%'
			outputText[outputIndex+1] = '2'
			outputText[outputIndex+2] = '0'
			outputIndex += 3
		} else {
			outputText[outputIndex] = text[i]
			outputIndex++
		}
	}

	return string(outputText)
}

func getSpacesCount(text string, start int, end int) int {
	counter := 0
	spaceCharacterCode := uint8(' ')

	for i := start; i < end; i++ {
		if text[i] == spaceCharacterCode {
			counter++
		}
	}

	return counter
}
