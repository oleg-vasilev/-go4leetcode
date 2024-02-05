package lc70

// #70 - Climbing stairs

// bruteforce recursion without dynamic programming
// heavy complexity - fails by time limit due redundant calculations
// func climbStairs(n int) int {
// 	if n <= 3 {
// 		return n
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

// straightforward Fibonacci
func climbStairs(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return y
}
