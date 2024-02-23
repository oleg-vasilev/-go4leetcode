package lc238

// 238. Product of Array Except Self

// Given an integer array nums, return an array answer such that answer[i]
// is equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Follow up: Can you solve the problem in O(1) extra space complexity?
// (The output array does not count as extra space for space complexity analysis.)

// straightforward bruteforce solution
// time:O(n*n) space:O(1)
// func productExceptSelf(nums []int) []int {
// 	var res = make([]int, len(nums))
//
// 	var cellRes int = 1
// 	for r, _ := range res {
// 		cellRes = 1
// 		for n, _ := range nums {
// 			if r != n {
// 				cellRes = cellRes * nums[n]
// 			}
// 		}
// 		res[r] = cellRes
// 	}
// 	return res
// }

// DP solution with precalculated multiplication parts
// time:O(n) space:O(n)
// possible improvement - use results array in-place instead one of dpLeft\dpRight
func productExceptSelf(nums []int) []int {
	var dpLeft = make([]int, len(nums))
	var dpRight = make([]int, len(nums))
	var res = make([]int, len(nums))

	var acc = 1
	for n := range nums {
		dpLeft[n] = acc
		acc = acc * nums[n]
	}
	acc = 1
	for n := len(nums) - 1; n >= 0; n-- {
		dpRight[n] = acc
		acc = acc * nums[n]
	}
	for r := range res {
		res[r] = dpLeft[r] * dpRight[r]
	}
	return res
}
