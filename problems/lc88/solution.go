package lc88

import (
	"slices"
)

// 88. Merge Sorted Array

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function,
// but instead be stored inside the array nums1. To accommodate this,
// nums1 has a length of m + n, where the first m elements denote the
// elements that should be merged, and the last n elements are
// set to 0 and should be ignored. nums2 has a length of n.

// dumbest solution ever xDD
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	slices.Sort(nums1)
	
}

// intuitive solution iterating from biggest lowest
func merge2(nums1 []int, m int, nums2 []int, n int) {
	mi := m - 1
	ni := n - 1
	
	// iterating over big slice from the end
	for i := m + n - 1; i >= 0; i-- {
		if ni < 0 {
			break
		}
		if mi >= 0 && nums1[mi] > nums2[ni] {
			nums1[i] = nums1[mi]
			mi--
		} else {
			nums1[i] = nums2[ni]
			ni--
		}
	}
}

// optimized but pretty complex
func merge3(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m == 0 || nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}
