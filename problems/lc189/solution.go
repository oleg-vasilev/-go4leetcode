package lc189

// 189 - Rotate array
//
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
// Example 1:
//
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:
//
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]
//
//
// Constraints:
//
// 1 <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1
// 0 <= k <= 105
//
// Follow up:
//
// Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
// Could you do it in-place with O(1) extra space?

func rotate(nums []int, k int) {
	rotateV4(nums, k)
}

// rotateV1 with simplest straightforward solution complexity: O(n), mem: O(n)
func rotateV1(nums []int, k int) {
	placements := make([]int, len(nums))
	for idx, val := range nums {
		newPos := (idx + k) % len(nums)
		placements[newPos] = val
	}
	copy(nums, placements)
}

// rotateV2 with slicing, appending and unpacking complexity: O(n), mem: O(n)
func rotateV2(nums []int, k int) {
	ln := len(nums)
	k %= ln
	copy(nums, append(nums[ln-k:], nums[:ln-k]...))
}

// rotateV3 with simple reverse array logic tricky applied three times complexity: O(n), mem: O(1)
func rotateV3(nums []int, k int) {
	reverse := func(nums []int, start int, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	ln := len(nums)
	k %= ln
	reverse(nums, 0, ln-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, ln-1)
}

func rotateV4(nums []int, k int) {
	numsSize := len(nums)
	k %= numsSize
	temp := make([]int, k)
	iterCount := 0
	// iterating numbers from last item to first
	for numID := numsSize - 1; numID >= 0; numID-- {
		iterCount = numsSize - 1 - numID
		// put value from current position into []temp if value will be out of range due to shift
		if iterCount < k {
			temp[k-1-iterCount] = nums[numID]
		}
		// fill the current position from the k-th element on the left if possible, or from []temp otherwise
		if numID-k >= 0 {
			nums[numID] = nums[numID-k]
		} else {
			nums[numID] = temp[numID]
		}
	}
}
