package lc27

// 27. Remove Element

// Given an integer array nums and an integer val, remove all occurrences
// of val in nums in-place. The order of the elements may be changed.
// Then return the number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal
// to val be k, to get accepted, you need to do the following things:
//
// Change the array nums such that the first k elements of nums
// contain the elements which are not equal to val. The remaining
// elements of nums are not important as well as the size of nums.
// Return k.

func removeElement(nums []int, val int) int {
	var res = make([]int, 0, len(nums))
	for _, v := range nums {
		if v != val {
			res = append(res, v)
		}
	}
	copy(nums, res)
	return len(res)
}

func removeElement2(nums []int, val int) int {
	counter := 0
	for _, v := range nums {
		if v != val {
			nums[counter] = v
			counter++
		}
	}
	return counter
}
