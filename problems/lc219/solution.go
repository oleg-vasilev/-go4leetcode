package lc219

// 219. Contains Duplicate II

// Given an integer array nums and an integer k,
// return true if there are two distinct indices i and j in the array
// such that nums[i] == nums[j] and abs(i - j) <= k.
//

// sliding window
// slow but memory-effective
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	if len(nums) < 2 || k == 0 {
// 		return false
// 	}
// 	var l, r = 0, 1
// 	for l < r-1 || r < len(nums) {
// 		if nums[l] == nums[r] {
// 			return true
// 		}
// 		if r-l < k && r < len(nums)-1 {
// 			r++
// 		} else {
// 			l++
// 			r = l + 1
// 		}
// 	}
// 	return false
// }

// same sliding window but refactored
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	for l := range nums {
// 		for r := l + 1; r <= l+k && r < len(nums); r++ {
// 			if nums[r] == nums[l] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// with item indexes memo - time effective but with extra for cache
func containsNearbyDuplicate(nums []int, k int) bool {
	var indexes = make(map[int]int)
	for idx, num := range nums {
		if storedIdx, ok := indexes[num]; ok && idx-storedIdx <= k {
			return true
		}
		indexes[num] = idx
	}
	return false
}
