package lc9

// 9. Palindrome Number

// Given an integer x, return true if x is a palindrome, and false otherwise.

// Follow up: Could you solve it without converting the integer to a string?

func isPalindrome(x int) bool {
	// easiest way
	// xs := strconv.Itoa(x)
	// runes := []rune(xs)
	// slices.Reverse(runes)
	// x2, _ := strconv.Atoi(string(runes))
	// return x == x2

	// negative not a palindrome
	if x < 0 {
		return false
	}

	// less than zero - always palindrome
	if x < 10 {
		return true
	}

	// ends to zero - not palindrome
	if x%10 == 0 {
		return false
	}

	// some 10-base math to extract digits from x one by one from the right
	// and add to reversed from the right, to build reversed number
	var reversed int

	// full reverse
	// tmp := x
	// for tmp != 0 {
	// 	reversed = 10*reversed + tmp%10
	// 	tmp /= 10
	// }
	// return reversed == x

	// instead of reversing whole number we can reverse only half and compare resulting numbers
	for x > 0 {
		reversed = reversed*10 + x%10 // 'reversed*10' will move current reversed digits to the left, and '+x%10' adds last digit from x to reversed
		x /= 10                       // removes last digit from X

		// checks for odd and even number of digits in initial number
		if reversed == x || reversed == x/10 {
			return true
		}
	}

	return false
}
