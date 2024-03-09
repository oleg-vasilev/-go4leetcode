package lc14

// 14. Longest Common Prefix

// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//

// func longestCommonPrefix(strs []string) string {
// 	// for starters assume that whole first item is common prefix
// 	var pref = strs[0]
// 	// iterate over other strings
// 	for _, str := range strs[1:] {
// 		// check longest common prefix
// 		i := 0
// 		for ; i < len(str) && i < len(pref) && pref[i] == str[i]; i++ {
// 			// do nothing
// 		}
// 		// cut prefix to common
// 		pref = pref[:i]
// 	}
// 	return pref
// }

// Using Prefix trie
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	// build trie and insert all strings
	t := &Trie{root: NewNode("")}
	for _, str := range strs {
		t.Insert(str)
	}

	var prefix string
	curr := t.root
	// iterate trie until node have only one child and this child isn't last in chain
	for t.CountChildren(curr) == 1 && !curr.isLast {
		for _, child := range curr.childs {
			if child != nil {
				// dive into found single non-nil child
				curr = child
				// increment prefix
				prefix += curr.value
				break
			}
		}
	}
	return prefix
}
