package lc125

import (
	"unicode"
)

// 125. Valid Palindrome

// A phrase is a palindrome if, after converting all uppercase letters
// into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters
// include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//

func isPalindrome(s string) bool {
	var runes = []rune(s)
	var leftPos int = 0
	var rightPos int = len(runes) - 1
	var leftRune, rightRune rune

	for leftPos <= rightPos {
		leftRune = unicode.ToLower(runes[leftPos])
		rightRune = unicode.ToLower(runes[rightPos])

		// check if left should be skipped
		if !isAlphanumeric(leftRune) {
			leftPos++
			continue
		}
		// check if right should be skipped
		if !isAlphanumeric(rightRune) {
			rightPos--
			continue
		}

		// runes should be compared
		if leftRune != rightRune {
			return false
		} else {
			leftPos++
			rightPos--
			continue
		}

	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
