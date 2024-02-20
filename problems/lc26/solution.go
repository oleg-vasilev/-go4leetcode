package lc26

import (
	"slices"
)

// 26. Remove Duplicates from Sorted Array

// Given an integer array nums sorted in non-decreasing order, remove
// the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same.
// Then return the number of unique elements in nums.
//
// Consider the number of unique elements of nums to be k, to get
// accepted, you need to do the following things:
//
// Change the array nums such that the first k elements of nums
// contain the unique elements in the order they were present in nums
// initially. The remaining elements of nums are not important
// as well as the size of nums.
// Return k.

// straightforward
func removeDuplicates(nums []int) int {
	var found = make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := found[v]; !ok {
			found[v] = struct{}{}
		}
	}
	var res = make([]int, 0, len(nums))
	for k, _ := range found {
		res = append(res, k)
	}
	slices.Sort(res)
	copy(nums, res)
	return len(found)
}

// more optimized
func removeDuplicatesOtp(nums []int) int {
	const tooBig = 999 // bigger than possible nums value
	var counter = len(nums)
	for i, _ := range nums {
		if i+1 >= len(nums) {
			continue
		}
		// check that current item same as next
		if nums[i] == nums[i+1] {
			// replace current
			nums[i] = tooBig
			counter--
		}
	}
	slices.Sort(nums)
	return counter
}
