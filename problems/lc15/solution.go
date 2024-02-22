package lc115

import (
	"fmt"
	"slices"
)

// 15. 3Sum

// Given an integer array nums, return all the triplets
// [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k,
// and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.

// straightforward bruteforce solution with bad time complexity
func threeSumBad(nums []int) [][]int {
	var (
		results = make([][]int, 0)
		resSet  = make(map[string]struct{})
		lastIdx = len(nums) - 1
	)
	for n1 := 0; n1 <= lastIdx-2; n1++ { // picked up first number

		for n2 := n1 + 1; n2 <= lastIdx-1; n2++ { // picked up second number
			for n3 := n2 + 1; n3 <= lastIdx; n3++ { // picked up third number
				// check combination legality
				if nums[n1]+nums[n2]+nums[n3] == 0 {
					// add to results
					triplet := []int{nums[n1], nums[n2], nums[n3]}
					slices.Sort(triplet)
					tripletStr := fmt.Sprintf("%v", triplet)
					if _, ok := resSet[tripletStr]; ok {
						continue
					} else {
						resSet[tripletStr] = struct{}{}
						results = append(results, triplet)
					}
				}
			}
		}

	}

	return results
}

func threeSum(nums []int) [][]int {
	var (
		results = make([][]int, 0)
		// resSet  = make(map[string]struct{})
		sum int
	)
	slices.Sort(nums)

	for n1 := 0; n1 < len(nums)-2; n1++ { // picked up first number
		// optimization - skipping identical values
		if n1 > 0 && nums[n1] == nums[n1-1] {
			continue
		}
		for n2, n3 := n1+1, len(nums)-1; n2 < n3; {
			sum = nums[n2] + nums[n3] + nums[n1]
			if sum == 0 {
				triplet := []int{nums[n1], nums[n2], nums[n3]}
				// memorizing discovered results to skip duplicates not needed
				// due to sorted input and skipping optimizations
				// slices.Sort(triplet)
				// tripletStr := fmt.Sprintf("%v", triplet)
				// if _, ok := resSet[tripletStr]; !ok {
				// 	resSet[tripletStr] = struct{}{}
				results = append(results, triplet)
				// }
				n2++
				n3--
				// optimization - skipping identical values
				for n2 < n3 && nums[n2] == nums[n2-1] {
					n2++
				}
				for n2 < n3 && nums[n3] == nums[n3+1] {
					n3--
				}
			} else if sum > 0 {
				n3--
			} else {
				n2++
			}
		}
	}
	return results
}
