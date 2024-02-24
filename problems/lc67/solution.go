package lc67

import (
	"slices"
	"strconv"
)

// 67. Add Binary

// Given two binary strings a and b, return their sum as a binary string.

// lazy implementation won't work on long inputs
// func addBinary(a string, b string) string {
// 	an, _ := strconv.ParseInt(a, 2, 64)
// 	bn, _ := strconv.ParseInt(b, 2, 64)
// 	return strconv.FormatInt(an+bn, 2)
// }

// simple lazy solution won't work with big inputs
// func addBinary(a string, b string) string {
// 	an, _ := strconv.ParseInt(a, 2, 64)
// 	bn, _ := strconv.ParseInt(b, 2, 64)
// 	return strconv.FormatInt(an+bn, 2)
// }
//
// // although this will work with larger numbers xD
// func addBinary(a string, b string) string {
// 	an, _ := new(big.Int).SetString(a, 2)
// 	bn, _ := new(big.Int).SetString(b, 2)
// 	return an.Add(an, bn).Text(2)
// }

// but it’s obvious that they want to see the implementation of binary addition
func addBinary(a string, b string) string {
	var overflow int
	var result = make([]byte, 0, len(a))

	var ai = len(a) - 1
	var bi = len(b) - 1

	for ai >= 0 || bi >= 0 {
		sum := overflow
		if ai >= 0 {
			sum += int(a[ai] - '0')
		}
		if bi >= 0 {
			sum += int(b[bi] - '0')
		}

		result = strconv.AppendInt(result, int64(sum%2), 2)
		overflow = sum / 2

		ai--
		bi--
	}
	if overflow == 1 {
		result = strconv.AppendInt(result, int64(overflow), 2)
	}
	slices.Reverse(result)
	return string(result)
}
