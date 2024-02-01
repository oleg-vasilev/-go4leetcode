package lc3

// Longest substring

func lengthOfLongestSubstring(s string) int {
	stringRunes := []rune(s)
	windowRunes := make(map[rune]struct{})
	var maxSubstringLen int
	var windowStart int
	for windowEnd := 0; windowEnd < len(stringRunes); windowEnd++ {
		
		// rune already seen
		if _, ok := windowRunes[stringRunes[windowEnd]]; ok {
			// update max length if current window is bigger than previous max
			if len(windowRunes) > maxSubstringLen {
				maxSubstringLen = len(windowRunes)
			}
			// cut the window runes one by one from the start to seen rune
			for stringRunes[windowStart] != stringRunes[windowEnd] {
				delete(windowRunes, stringRunes[windowStart])
				windowStart++
			}
			// cut the seen rune and mark new window start
			delete(windowRunes, stringRunes[windowStart])
			windowStart++
		}
		// add current end rune to window
		windowRunes[stringRunes[windowEnd]] = struct{}{}
	}
	// last known window can be bigger than previous max
	if len(windowRunes) > maxSubstringLen {
		maxSubstringLen = len(windowRunes)
	}
	return maxSubstringLen
}
