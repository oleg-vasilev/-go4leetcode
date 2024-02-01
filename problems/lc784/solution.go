// Package lc784 - Letter case permutations
package lc784

import (
	"unicode"
)

// straightforward DFS-like solution

func letterCasePermutation(source string) []string {
	res := make([]string, 0)
	casePermutations(&res, source, "")
	return res
}

func casePermutations(res *[]string, source string, perm string) {
	if len(perm) == len(source) {
		*res = append(*res, perm)
		return
	}
	r := rune(source[len(perm)])
	if unicode.IsLetter(r) {
		casePermutations(res, source, perm+string(unicode.ToLower(r)))
		casePermutations(res, source, perm+string(unicode.ToUpper(r)))
	} else {
		casePermutations(res, source, perm+string(r))
	}
}
