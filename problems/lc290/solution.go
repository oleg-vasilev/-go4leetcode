package lc290

import (
	"strings"
)

// 290. Word Pattern

// Given a pattern and a string s, find if s follows the same pattern.
//
// Here follow means a full match, such that there is a bijection between
// a letter in pattern and a non-empty word in s.

func wordPattern(pattern string, s string) bool {
	var matches = make(map[rune]string)
	var matchesBack = make(map[string]rune)
	var pt = []rune(pattern)
	var split = strings.Split(s, " ")

	if len(pt) != len(split) {
		return false
	}

	for i := 0; i < len(pt); i++ {
		if v, ok := matches[pt[i]]; ok {
			if v != split[i] {
				return false
			}
		} else {
			if _, ok := matchesBack[split[i]]; ok {
				return false
			}
			matches[pt[i]] = split[i]
			matchesBack[split[i]] = pt[i]
		}
	}
	return true
}
