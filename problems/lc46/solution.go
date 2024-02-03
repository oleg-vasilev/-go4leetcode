package lc46

// 46 - Permutations

// straightforward recursive solution without classic backtracking
func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	result := make([][]int, 0)
	currItem := nums[0]
	permsWithoutCurr := permute(nums[1:])
	for _, permWithoutCurr := range permsWithoutCurr {
		for pos := 0; pos <= len(permWithoutCurr); pos++ {
			perm := make([]int, 0)
			perm = append(perm, permWithoutCurr[:pos]...)
			perm = append(perm, currItem)
			perm = append(perm, permWithoutCurr[pos:]...)
			result = append(result, perm)
		}
	}
	return result
}

// classic backtracking solution (tracking used items with map)

// func permute(nums []int) [][]int {
// 	result := make([][]int, 0)
// 	backtrackPermutations(&result, nums, []int{}, map[int]struct{}{})
// 	return result
// }
//
// func backtrackPermutations(result *[][]int, items []int, perm []int, used map[int]struct{}) {
// 	// target reached - permutation length match required
// 	if len(perm) == len(items) {
// 		res := make([]int, len(perm))
// 		copy(res, perm)
// 		*result = append(*result, res)
// 		return
// 	}
// 	// pick items from numbers to construct new permutation
// 	for i := 0; i < len(items); i++ {
// 		// skip items already used in current permutation branch
// 		if _, ok := used[items[i]]; ok {
// 			continue
// 		}
// 		// made available choice
// 		used[items[i]] = struct{}{}
// 		perm = append(perm, items[i])
// 		// calculate branch with the choice made above
// 		backtrackPermutations(result, items, perm, used)
// 		// undo last made choice, to prepare state for next iteration and alternative choices
// 		delete(used, items[i])
// 		perm = perm[:len(perm)-1]
// 	}
// }

// backtracking solution with swaps inside original array, without tracking used items
// func permute(nums []int) [][]int {
// 	result := make([][]int, 0)
// 	findPermutations(&result, nums, 0)
// 	return result
// }
//
// func findPermutations(result *[][]int, items []int, index int) {
// 	// current index out of bounds - current permutation is complete - add it to results
// 	if index == len(items) {
// 		res := make([]int, len(items))
// 		copy(res, items)
// 		*result = append(*result, res)
// 		return
// 	}
//
// 	// iterate over items and swap different items in pairs to create new permutations
// 	for i := index; i < len(items); i++ {
// 		if index == i {
// 			// can't swap item with itself - increase index and proceed
// 			findPermutations(result, items, index+1)
// 		} else {
// 			// swap items to create new permutation, increase index and proceed.
// 			items[i], items[index] = items[index], items[i]
// 			findPermutations(result, items, index+1)
// 			// swap items back to cancel last permutation and try next one on next loop
// 			items[i], items[index] = items[index], items[i]
// 		}
// 	}
// }
