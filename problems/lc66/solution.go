package lc66

// 66. Plus One

// You are given a large integer represented as an integer array digits,
// where each digits[i] is the ith digit of the integer.
// The digits are ordered from most significant to least significant
// in left-to-right order. The large integer does not contain any leading 0's.
//
// Increment the large integer by one and return the resulting array of digits.

func plusOne(digits []int) []int {
	var tmp int
	var pos = len(digits) - 1
	digits[pos]++
	for pos >= 0 {
		digits[pos] += tmp
		if digits[pos] >= 10 {
			tmp = digits[pos] / 10
			digits[pos] %= 10
			pos--
		} else {
			return digits
		}
	}
	var res = make([]int, 0, len(digits)+1)
	if tmp > 0 {
		res = append(res, tmp)
		res = append(res, digits...)
		return res
	} else {
		return digits
	}
}
