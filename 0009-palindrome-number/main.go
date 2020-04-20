package main

func isPalindrome(x int) bool {
	if x < 0 {
		// negatives are never palindromes
		return false
	}

	xx := x
	// same logic as reverse integer problem
	var reversed int
	for ; xx != 0; xx /= 10 {
		reversed *= 10
		reversed += xx - ((xx / 10) * 10)
	}
	return reversed == x
}
