package lc784

// backtracking zero-imports solution

func letterCasePermutationBacktrack(s string) []string {
	var res []string
	var curr []byte
	backtrack(s, 0, curr, &res)
	return res
}

func backtrack(s string, start int, curr []byte, res *[]string) {
	if len(curr) == len(s) {
		*res = append(*res, string(curr))
		return
	}
	if isAlpha(s[start]) {
		curr = append(curr, toLower(s[start]))
		backtrack(s, start+1, curr, res)
		curr = curr[:len(curr)-1]
		curr = append(curr, toUpper(s[start]))
		backtrack(s, start+1, curr, res)
		// curr = curr[:len(curr)-1]
	} else {
		curr = append(curr, s[start])
		backtrack(s, start+1, curr, res)
		// curr = curr[:len(curr)-1]
	}
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return 'a' + ch - 'A'
	}
	return ch
}

func toUpper(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		return 'A' + ch - 'a'
	}
	return ch
}
