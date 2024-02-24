package lc242

import (
	"slices"
)

// 242. Valid Anagram

// Given two strings s and t, return true if t is an anagram of s,
// and false otherwise.
//
// An Anagram is a word or phrase formed by rearranging the letters
// of a different word or phrase, typically using all the original
// letters exactly once.

func isAnagram(s string, t string) bool {
	var sr = []rune(s)
	var tr = []rune(t)
	var table = make(map[rune]int)

	for _, r := range sr {
		if count, ok := table[r]; ok {
			table[r] = count + 1
		} else {
			table[r] = 1
		}
	}

	for _, r := range tr {
		count, ok := table[r]
		if ok {
			count--
			if count == 0 {
				delete(table, r)
			} else {
				table[r] = count
			}
		} else {
			return false
		}
	}
	return len(table) == 0
}

// using std lib
func isAnagram2(s string, t string) bool {
	var sr = []rune(s)
	var st = []rune(t)
	slices.Sort(sr)
	slices.Sort(st)
	return slices.Equal(sr, st)
}
