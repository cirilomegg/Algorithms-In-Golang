package compression

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompress(t *testing.T) {
	data := map[string]string{}
	data["aaabccccdd"] = "a3b1c4d2"
	data["abbcccccdee"] = "a1b2c5d1e2"
	data["abcd"] = "abcd"
	data["aaa"] = "a3"

	for key, value := range data {
		compressed := Compress(key)
		assert.Equal(t, value, compressed)
	}
}
