package lc77

/**

#77 - Combinations

Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
You may return the answer in any order


Example 1:

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
Example 2:

Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.

*/

/** iterative non-backtracking answer based on the fact that initial items is just ascending sequence from 1 to n.

1 Create an array of k length and start at index 0.
2 Increment the value at the current index i.
3 If the value at the current index i has exceeded the limit n go back to the previous index i-- and continue incrementing for the next combinations.
4 If the last index k - 1 has been reached, append a copy of the current array arr as it is a completed combination.
5 If the value at the current index has not exceeded the limit n and the last index k - 1 has not been reached, move to the next index i++ and copy the value from the previous index arr[i] = arr[i - 1] then continue from step 2.
*/
func combine(n int, k int) [][]int {
	var res [][]int
	buf := make([]int, k)
	for i := 0; i >= 0; {
		buf[i]++
		if buf[i] > n {
			i--
		} else if i == k-1 { // last place filled - combination is ready
			comb := make([]int, k)
			copy(comb, buf)
			res = append(res, comb)
		} else {
			i++
			buf[i] = buf[i-1]
		}
	}
	return res
}

// Recursive backtracking solution
func combineBacktrack(n int, k int) [][]int {
	var results [][]int
	current := make([]int, k)
	
	var backtrack func(int)
	
	backtrack = func(idx int) {
		if idx == k {
			result := make([]int, k)
			copy(result, current)
			results = append(results, result)
			return
		}
		
		startIdx := 1
		if idx > 0 {
			startIdx = current[idx-1] + 1
		}
		
		for i := startIdx; i <= n; i++ {
			current[idx] = i
			backtrack(idx + 1)
		}
	}
	
	backtrack(0)
	
	return results
}
