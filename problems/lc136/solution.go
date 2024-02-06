package lc136

// #136 - Single Number
// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.

func singleNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
