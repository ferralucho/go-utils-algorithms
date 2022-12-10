package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	assert.Equal(t, "", longestSubstringWithoutRepeatingCharacters(""))

	assert.Equal(t, "abc", longestSubstringWithoutRepeatingCharacters("abc"))

	assert.Equal(t, "abc", longestSubstringWithoutRepeatingCharacters("abcabcbb"))

	assert.Equal(t, "a", longestSubstringWithoutRepeatingCharacters("aaaaaa"))
}
