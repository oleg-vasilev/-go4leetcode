package lc137

// 137. Single Number II

// Given an integer array nums where every element appears three
// times except for one, which appears exactly once.
// Find the single element and return it.
//
// You must implement a solution with a linear runtime complexity
// and use only constant extra space.
//

func singleNumber(nums []int) int {
	var first, second int // bit appearances
	for _, n := range nums {
		first ^= n & ^second
		second ^= n & ^first
	}
	return first
}
