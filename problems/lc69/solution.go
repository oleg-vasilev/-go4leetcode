package lc69

// 69. Sqrt(x)

// Given a non-negative integer x, return the square root of x rounded down
// to the nearest integer. The returned integer should be non-negative as well.
//
// You must not use any built-in exponent function or operator.

func mySqrt(x int) int {
	var low int
	var high = x
	var mid int
	for low <= high {
		mid = (low + high) / 2
		// mid = low + (high-low)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}
