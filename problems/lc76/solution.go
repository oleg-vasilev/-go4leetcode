package lc76

// 76. Minimum Window Substring

// Given two strings s and t of lengths m and n respectively,
// return the minimum window substring of s such that every
// character in t (including duplicates) is included in the window.
// If there is no such substring, return the empty string "".
//
// The testcases will be generated such that the answer is unique.

// straightforward solution
func minWindow(s string, t string) string {
	// handle obvious case
	if len(s) < len(t) {
		return ""
	}
	// build default frequency map of t
	var needs = Needs{}
	for _, r := range []rune(t) {
		count := needs[r]
		needs[r] = count + 1
	}

	var result = Result{}
	var runes = []rune(s)
	var l, r = 0, -1
	for {
		// check if our needs are satisfied
		if needs.Satisfied() {
			// add current indexes as possible result
			result.Add(l, r)
			// shrink window from left side and go to next iteration:
			if v, ok := needs[runes[l]]; ok {
				// left rune was in our needs list - so we will need it again
				needs[runes[l]] = v + 1
			}
			l++
			continue
		}
		// we are not satisfied ...
		if r < len(runes)-1 {
			// .. but we not at the end, so expand window to the right
			r++
			if v, ok := needs[runes[r]]; ok {
				// new rune on the right was is in our needs list
				needs[runes[r]] = v - 1
			}

		} else {
			// .. and we are at the end - exit loop
			break
		}
	}
	return result.Compute(runes)
}

type Needs map[rune]int

func (n Needs) Satisfied() bool {
	for _, v := range n {
		if v > 0 {
			return false
		}
	}
	return true
}

// Result can store all possible combinations if needed
// but for now it stores only one result with minimal length
type Result struct {
	left, right int
	found       bool
}

func (r *Result) Add(left, right int) {
	if !r.found {
		r.left = left
		r.right = right
		r.found = true
	} else {
		if r.right-r.left > right-left {
			r.left = left
			r.right = right
		}
	}
}
func (r *Result) Compute(runes []rune) string {
	if r.found {
		if r.right == len(runes)-1 {
			return string(runes[r.left:])
		} else {
			return string(runes[r.left : r.right+1])
		}
	} else {
		return ""
	}
}
