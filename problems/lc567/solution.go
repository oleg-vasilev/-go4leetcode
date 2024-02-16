package lc567

// 567. Permutation in String

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
//
// In other words, return true if one of s1's permutations is the substring of s2.

// Example 1:
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").

// Example 2:
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	exampleSet := newExampleSet([]byte(s1))
	runes := []byte(s2)
	windowSet := make(map[byte]int)
	windowStart := 0
	windowEnd := 0
	for windowEnd < len(s2) {
		// not satisfies s1
		if _, ok := exampleSet[runes[windowEnd]]; !ok {
			// reset current window set
			windowSet = make(map[byte]int)
			// set window start position to next item
			windowEnd++
			windowStart = windowEnd
			continue
		}
		// move window end
		windowSet[runes[windowEnd]]++
		windowEnd++
		// move window start
		if windowEnd-windowStart > len(s1) {
			if count, ok := windowSet[runes[windowStart]]; ok && count > 1 {
				windowSet[runes[windowStart]]--
			} else {
				delete(windowSet, runes[windowStart])
			}
			windowStart++
		}
		if mapsEquals(exampleSet, windowSet) {
			return true
		}
	}
	return false
}

func newExampleSet(s []byte) map[byte]int {
	set := make(map[byte]int)
	for _, r := range s {
		set[r]++
	}
	return set
}

func mapsEquals(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if w, ok := m2[k]; !ok || v != w {
			return false
		}
	}
	return true
}

// optimized for constraints (s1 and s2 consist of lowercase English letters.)
func checkInclusionOptimized(s1 string, s2 string) bool {
	need := [26]int{} // position encodes letter, 0 for a, 25 for z, value - letter count in string
	for i := range s1 {
		need[s1[i]-'a']++
	}
	
	have := [26]int{}
	for r := range s2 {
		if r >= len(s1) {
			have[s2[r-len(s1)]-'a']--
		}
		have[s2[r]-'a']++
		if need == have {
			return true
		}
	}
	return false
}
