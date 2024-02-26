package lc209

// 209. Minimum Size Subarray Sum

// Given an array of positive integers nums and a positive integer target,
// return the minimal length of a subarray whose sum is greater than
// or equal to target. If there is no such subarray, return 0 instead.

// straightforward solution with both concepts - prefixSum and sliding window
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 1
		} else {
			return 0
		}
	}
	// build standard prefix sum array
	var prefSum = make([]int, len(nums))
	for i, n := range nums {
		// and check one-length instant answers at the same time
		if n == target {
			return 1
		}
		prefSum[i] = n
		if i > 0 {
			prefSum[i] += prefSum[i-1]
		}
	}
	// do sliding window over prefix sum array
	var minWindowLength int // tracking min window length when window sum satisfies target
	var windowSum int       // just pre-allocated memory for calc window sum on each iteration
	// star with smallest window from the left
	for l, r := 0, 1; l < r && r < len(nums); {
		// calc window sum for current window
		windowSum = prefSum[r]
		if l > 0 {
			windowSum -= prefSum[l-1]
		}
		// handle case when window sum is not satisfies target
		if windowSum < target {
			// knowing that we do not have negative numbers
			// expanding window always leads to increased sum
			// expand window to right and continue to next iteration
			r++
			continue
		}
		// window sum satisfies target - update min length we have
		if minWindowLength == 0 {
			minWindowLength = r - l + 1
		} else {
			minWindowLength = min(minWindowLength, r-l+1)
		}
		// and try to shrink window from the left, so next iteration will
		// be able to check if sum of smaller window still satisfies target
		l++
	}
	return minWindowLength
}
