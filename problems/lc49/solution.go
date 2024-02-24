package lc49

import (
	"slices"
)

// 49. Group Anagrams

// Given an array of strings strs, group the anagrams together.
// You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters
// of a different word or phrase, typically using all the original
// letters exactly once.

func groupAnagrams(strs []string) [][]string {
	var groups = make(map[string][]string)

	for _, s := range strs {
		rn := []rune(s)
		slices.Sort(rn)
		normal := string(rn)
		if gr, ok := groups[normal]; ok {
			gr = append(gr, s)
			groups[normal] = gr
		} else {
			groups[normal] = []string{s}
		}
	}

	var result = make([][]string, 0)

	for _, v := range groups {
		result = append(result, v)
	}
	return result
}
