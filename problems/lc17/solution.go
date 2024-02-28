package lc17

// 17. Letter Combinations of a Phone Number

// Given a string containing digits from 2-9 inclusive, return all possible
// letter combinations that the number could represent. Return the answer in any order.
//
// A mapping of digits to letters (just like on the telephone buttons)
// is given below. Note that 1 does not map to any letters.

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var results = make([]string, 0)

	var digitsMap = map[uint8][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var backtrack func(comb string, digits string)
	backtrack = func(comb string, digits string) {
		if len(digits) == 0 {
			results = append(results, comb)
			return
		}
		letters, ok := digitsMap[digits[0]]
		if !ok {
			panic("invalid digit input")
		}
		digitsLeft := digits[1:]
		for _, l := range letters {
			backtrack(comb+string(l), digitsLeft)
		}
	}
	backtrack("", digits)
	return results
}
