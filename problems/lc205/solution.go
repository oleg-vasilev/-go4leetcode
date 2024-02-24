package lc205

// 205. Isomorphic Strings

// Given two strings s and t, determine if they are isomorphic.
//
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
//
// All occurrences of a character must be replaced with another character
// while preserving the order of characters. No two characters may map to
// the same character, but a character may map to itself.

func isIsomorphic(s string, t string) bool {
	var sr, tr = []rune(s), []rune(t)
	var s2t, t2s = make(map[rune]rune), make(map[rune]rune)
	for i := 0; i < len(s); i++ {
		// check if sr[i] already mapped to something
		if m, ok := s2t[sr[i]]; ok {
			if m != tr[i] {
				// sr[i] mapped buy runes does not match
				return false
			}
		} else {
			// check if we can map sr[i] to tr[i]
			if _, ok := t2s[tr[i]]; ok {
				return false // tr[i] already mapped into something from sr
			} else {
				s2t[sr[i]] = tr[i]
				t2s[tr[i]] = sr[i]
			}
		}
	}
	return true
}
