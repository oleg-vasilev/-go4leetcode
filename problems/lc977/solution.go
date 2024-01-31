// Package lc977 - Sorted Squares
package lc977

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l := 0
	r := len(nums) - 1
	ls := nums[l] * nums[l]
	rs := nums[r] * nums[r]
	pos := len(nums) - 1
	for pos >= 0 {
		if ls > rs {
			res[pos] = ls
			if l < len(nums) {
				l++
				ls = nums[l] * nums[l]
			}
		} else {
			res[pos] = rs
			if r > 0 {
				r--
				rs = nums[r] * nums[r]
			}
		}
		pos--
	}
	return res
}
