package lc35

// 35 - search insert position

func searchInsert(nums []int, target int) int {
	var first, last, mid int
	last = len(nums) - 1
	for {
		mid = (first + last) / 2
		if first > last {
			return first
		}
		if last < first {
			return last
		}
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			last = mid - 1
		} else if target > nums[mid] {
			first = mid + 1
		}
	}
}
