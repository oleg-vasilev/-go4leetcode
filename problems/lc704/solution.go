package binsearch

// lc704 Binary search

func search(nums []int, target int) int {
	var first, last, mid int
	last = len(nums) - 1
	for {
		mid = (first + last) / 2
		if first > last {
			return -1
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
}
