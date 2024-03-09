package lc151

import (
	"slices"
	"strings"
)

// 151. Reverse Words in a String

// Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters.
// The words in s will be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
// Note that s may contain leading or trailing spaces or multiple
// spaces between two words. The returned string should only have
// a single space separating the words. Do not include any extra spaces.
//

func reverseWords(s string) string {
	var parts = strings.Split(strings.TrimSpace(s), " ") // actually strings.TrimSpace is optional, slices.DeleteFunc will remove empty items anyway
	slices.Reverse(parts)
	parts = slices.DeleteFunc(parts, func(item string) bool {
		return len(item) == 0
	})
	return strings.Join(parts, " ")
}
