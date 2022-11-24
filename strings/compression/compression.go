package compression

import (
	"strconv"
	"strings"
)

func Compress(text string) string {
	textLength := len(text)

	builder := strings.Builder{}

	characterCounter := 0

	for i := 0; i < textLength; i++ {
		characterCounter++

		if i+1 == textLength || text[i] != text[i+1] {
			builder.WriteString(string(text[i]))
			builder.WriteString(strconv.Itoa(characterCounter))
			characterCounter = 0
		}
	}

	if builder.Len() > textLength {
		return text
	}

	return builder.String()
}
