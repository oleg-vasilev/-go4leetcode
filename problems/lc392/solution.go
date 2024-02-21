package lc392

// 392. Is Subsequence

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//
// A subsequence of a string is a new string that is formed from the original
// string by deleting some (can be none) of the characters without disturbing
// the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//

func isSubsequence(s string, t string) bool {
	var sub = []rune(s)
	var sequence = []rune(t)

	var lookFrom int
	var nextMatchFound bool

	// iterate over runes of subsequence we looking for
	for _, v := range sub {
		nextMatchFound = false // reset flag no each matching loop iteration
		// look in sequence for rune 'v' starting from position next to previous match
		for j := lookFrom; j < len(sequence); j++ {
			// match found
			if sequence[j] == v {
				lookFrom = j + 1      // looking from next index next time
				nextMatchFound = true // to make sure we
				break
			}
		}
		if !nextMatchFound {
			return false
		}
	}
	return true
}
