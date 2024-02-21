package lc45

// 45. Jump Game 2

// You are given a 0-indexed array of integers nums of length n.
// You are initially positioned at nums[0].
//
// Each element nums[i] represents the maximum length of a forward
// jump from index i. In other words, if you are at nums[i],
// you can jump to any nums[i + j] where:
// 0 <= j <= nums[i] and i + j < n

// Return the minimum number of jumps to reach nums[n - 1].
// The test cases are generated such that you can reach nums[n - 1].

// greedy non-trivial - jump as far as possible every time time:O(n), space:O(1)
// func jump(nums []int) int {
// 	var jumpsCount int         // self-explanatory
// 	var curMaxReachIdx int     // holds currently max reachable position
// 	var maxPossibleJumpIdx int // holds max index we can reach by jumping from one already reachable cells
//
// 	// check cells one by one
// 	for i := 0; i < len(nums)-1; i++ {
//
// 		// check and update max jump in case if we jump from current cell
// 		if maxPossibleJumpIdx < i+nums[i] {
// 			maxPossibleJumpIdx = i + nums[i]
// 		}
//
// 		// all currently reachable cells checked - jump as far as we can
// 		if i == curMaxReachIdx {
// 			curMaxReachIdx = maxPossibleJumpIdx // update max reachable position
// 			maxPossibleJumpIdx = 0              // reset maxDistance
// 			jumpsCount++                        // count jump
// 		}
// 	}
// 	return jumpsCount
// }

// recursive + greedy combination
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	// just for readability
	lastIdx := len(nums) - 1
	// find next biggest jump from the end
	var idx int
	for i := lastIdx - 1; i >= 0; i-- {
		if nums[i] >= lastIdx-i {
			idx = i
		}
	}
	return 1 + jump(nums[:idx+1])
}
