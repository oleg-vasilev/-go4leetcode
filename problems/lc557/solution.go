package lc557

// 557. Reverse Words in a String III

// Given a string s, reverse the order of characters in each word within
// a sentence while still preserving whitespace and initial word order.

// Example 1:
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

// Example 2:
// Input: s = "Mr Ding"
// Output: "rM gniD"

func reverseWords(s string) string {
	bytes := []byte(s)
	wordStartIdx := -1
	wordLength := 0
	// iterate over slice of bytes detecting start and end of each word
	for i := 0; i < len(bytes); i++ {
		if bytes[i] != 0x20 {
			wordLength++
			if wordStartIdx == -1 {
				wordStartIdx = i
			}
			if i == len(bytes)-1 {
				reverseString(bytes, wordStartIdx, wordStartIdx+wordLength-1)
			}
		} else {
			reverseString(bytes, wordStartIdx, wordStartIdx+wordLength-1)
			wordStartIdx, wordLength = -1, 0
		}
	}
	return string(bytes)
}

func reverseString(s []byte, from, to int) {
	if from < 0 || to <= 0 || from >= to || to >= len(s) {
		return
	}
	for from < to {
		s[from], s[to] = s[to], s[from]
		from++
		to--
	}
}
