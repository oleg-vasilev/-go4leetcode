package lc55

// 55. Jump Game

// You are given an integer array nums. You are initially positioned
// at the array's first index, and each element in the array represents
// your maximum jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
//

// recursive bruteforce

// func canJump(nums []int) bool {
// 	jmp := nums[0]
// 	if jmp >= len(nums)-1 {
// 		return true
// 	}
// 	if jmp == 0 {
// 		return false
// 	}
//
// 	var result bool
// 	for j := 1; j <= jmp; j++ {
// 		result = result || canJump(nums[j:])
// 		// optimization #1
// 		if result {
// 			break
// 		}
// 	}
//
// 	return result
// }

// dp with memoization
func canJump(nums []int) bool {

	var memo = make(map[int]bool)
	var jumpCheck func(nums []int, memo map[int]bool) bool

	jumpCheck = func(nums []int, memo map[int]bool) bool {
		jmp := nums[0]
		if jmp >= len(nums)-1 {
			return true
		}
		if jmp == 0 {
			return false
		}
		if storedRes, ok := memo[len(nums)]; ok {
			return storedRes
		}

		var result bool
		for j := 1; j <= jmp; j++ {
			result = result || jumpCheck(nums[j:], memo)
			if result {
				break
			}
		}
		memo[len(nums)] = result
		return result
	}

	return jumpCheck(nums, memo)
}

// greedy
func canJumpGreedy(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	destinationIdx := len(nums) - 1
	// check that destination is reachable from previous cell
	for i := destinationIdx - 1; i >= 0; i-- {
		if i+nums[i] >= destinationIdx {
			// and if it is reachable - change destination to this cell
			destinationIdx = i
		}
	}
	return destinationIdx == 0 // if we passed to start original destination is reachable
}
