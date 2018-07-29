package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	textToSearch := ""
	text := ""
	assert.Equal(t, NotFound, Search(textToSearch, text))

	textToSearch = "a"
	text = ""
	assert.Equal(t, NotFound, Search(textToSearch, text))

	textToSearch = "ab"
	text = "c"
	assert.Equal(t, NotFound, Search(textToSearch, text))

	textToSearch = "abc"
	text = "c"
	assert.Equal(t, "3", Search(textToSearch, text))

	textToSearch = "abcabc"
	text = "ab"
	assert.Equal(t, "1, 4", Search(textToSearch, text))

	textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	text = "Peter"
	assert.Equal(t, "1, 20, 75", Search(textToSearch, text))

	textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	text = "peter"
	assert.Equal(t, "1, 20, 75", Search(textToSearch, text))

	textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	text = "pick"
	assert.Equal(t, "30, 58", Search(textToSearch, text))

	textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	text = "pi"
	assert.Equal(t, "30, 37, 43, 51, 58", Search(textToSearch, text))

	textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	text = "z"
	assert.Equal(t, NotFound, Search(textToSearch, text))

	textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	text = "Peterz"
	assert.Equal(t, NotFound, Search(textToSearch, text))
}

func TestToLowerCase(t *testing.T) {
	text := ""
	assert.Equal(t, "", ToLowerCase(text))

	text = "A"
	assert.Equal(t, "a", ToLowerCase(text))

	text = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	assert.Equal(t, "abcdefghijklmnopqrstuvwxyz", ToLowerCase(text))

	text = "ABC < D"
	assert.Equal(t, "abc < d", ToLowerCase(text))

	text = "AB 上海 d"
	assert.Equal(t, "ab 上海 d", ToLowerCase(text))
}
