package lc228

import (
	"strconv"
	"strings"
)

// 228. Summary Ranges

// You are given a sorted unique integer array nums.
//
// A range [a,b] is the set of all integers from a to b (inclusive).
//
// Return the smallest sorted list of ranges that cover all
// the numbers in the array exactly. That is, each element of nums
// is covered by exactly one of the ranges, and there is no
// integer x such that x is in one of the ranges but not in nums.
//
// Each range [a,b] in the list should be output as:
//
// "a->b" if a != b
// "a" if a == b

func summaryRanges(nums []int) []string {
	var results []string
	res := strings.Builder{}
	var left, right int
	for left < len(nums) {
		res.WriteString(strconv.Itoa(nums[left]))
		right = left + 1
		for right < len(nums) && nums[right]-nums[right-1] == 1 {
			right++
		}
		if right-1 > left {
			res.WriteString("->")
			res.WriteString(strconv.Itoa(nums[right-1]))
		}
		results = append(results, res.String())
		res.Reset()
		left = right
	}
	return results
}
