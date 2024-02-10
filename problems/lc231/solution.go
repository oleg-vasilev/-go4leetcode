package lc231

// #231 - Power of Two

// Given an integer n, return true if it is a power of two. Otherwise, return false.
// An integer n is a power of two, if there exists an integer x such that n == 2x.

// math
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// recursive
// func isPowerOfTwo(n int) bool {
// 	if n <= 0 {
// 		return false
// 	}
// 	if n == 1 {
// 		return true
// 	}
// 	if n%2 == 1 {
// 		return false
// 	}
// 	return isPowerOfTwo(n / 2)
// }
