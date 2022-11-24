package replace_spaces

import (
	"fmt"
	"strings"
	"testing"
)

func TestReplaceSpaces(t *testing.T) {
	url := "Some entry to replace      "
	result := ReplaceSpaces(url, len(strings.TrimSpace(url)))

	fmt.Println(result)

	result = strings.ReplaceAll(strings.TrimSpace(url), " ", "%20")
	fmt.Println(result)
}
