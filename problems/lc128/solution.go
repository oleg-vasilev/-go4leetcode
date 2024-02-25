package lc128

// 128. Longest Consecutive Sequence

// Given an unsorted array of integers nums,
// return the length of the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.

// sort with standard lib, actually not O(n) time (quick sort can be from O(n*logn) to O(n*n) in worst) but it works.

// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	slices.Sort(nums)
// 	var maxSeries, currentSeries int
// 	for i, _ := range nums {
// 		if i < 0 {
// 			continue
// 		}
// 		if nums[i] == nums[i-1] {
// 			continue
// 		} else if nums[i]-nums[i-1] == 1 {
// 			currentSeries++
// 		} else {
// 			maxSeries = max(currentSeries, maxSeries)
// 			currentSeries = 0
// 		}
// 	}
// 	return max(currentSeries, maxSeries) + 1
// }

func longestConsecutive(nums []int) int {
	// create and fill set of our numbers to be able to check existence for O(1)
	var set = make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	// sequence counter.
	var maxSeq, currSeq int
	for n := range set {

		// check if set contains previous item - and skip
		if _, ok := set[n-1]; ok {
			continue
		}
		// no prev item - so this is possible sequence start
		currSeq = 1
		for {
			if _, ok := set[n+currSeq]; ok {
				currSeq++
				continue
			}
			maxSeq = max(maxSeq, currSeq)
			break
		}
	}
	return maxSeq
}
