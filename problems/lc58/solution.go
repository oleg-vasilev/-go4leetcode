package lc58

import (
	"strings"
)

// 58. Length of Last Word

// Given a string s consisting of words and spaces,
// return the length of the last word in the string.
//
// A word is a maximal
// substring
// consisting of non-space characters only.

func lengthOfLastWord(s string) int {
	parts := strings.Split(strings.TrimSpace(s), " ")
	return len(parts[len(parts)-1])
}
