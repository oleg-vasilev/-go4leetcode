package lc14

// 14. Longest Common Prefix

func longestCommonPrefix(strs []string) string {
	// for starters assume that whole first item is common prefix
	var pref = strs[0]
	// iterate over other strings
	for _, str := range strs[1:] {
		// check longest common prefix
		i := 0
		for ; i < len(str) && i < len(pref) && pref[i] == str[i]; i++ {
			// do nothing
		}
		// cut prefix to common
		pref = pref[:i]
	}
	return pref
}
