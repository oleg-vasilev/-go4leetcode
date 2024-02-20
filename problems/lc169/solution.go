package lc169

import (
	"slices"
)

// 169. Majority Element

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

// Follow-up: Could you solve the problem in linear time and in O(1) space?

// straightforward
func majorityElement(nums []int) int {
	slices.Sort(nums)
	var majElem, majCount int
	majElem = nums[0]
	majCount = 1
	var counter int
	for i, _ := range nums {
		if i == 0 {
			counter++
			continue
		}
		if nums[i] == nums[i-1] {
			// current same as prev
			counter++
		} else {
			// current differs from prev
			// check if counter greater than found before
			if counter > majCount {
				majCount = counter
				majElem = nums[i-1]
			}
			// in any case reset counter to count new series
			counter = 1
		}
	}
	if counter > majCount {
		majCount = counter
		majElem = nums[len(nums)-1]
	}
	return majElem
}

// default rates calculation
func majorityElement2(nums []int) int {
	var rates = make(map[int]int)
	var item, rate int
	for _, v := range nums {
		r, ok := rates[v]
		if ok {
			rates[v] = r + 1
		} else {
			rates[v] = 1
		}
		if rates[v] > rate {
			rate = r
			item = v
		}
	}
	return item
}

// Boyer–Moore majority vote
func majorityElement3(nums []int) int {
	var element, counter int
	for _, v := range nums {
		if counter == 0 {
			element = v
		}
		if v == element {
			counter++
		} else {
			counter--
		}
	}
	return element
}
