package lc28

import (
	"strings"
)

// 28. Find the Index of the First Occurrence in a String

// Given two strings needle and haystack, return the index of the first
// occurrence of needle in haystack, or -1 if needle is not part of haystack.

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
