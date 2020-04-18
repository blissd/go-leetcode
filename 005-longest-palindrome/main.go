package main

func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	// each element represents the length of the palindrome at that index in the string
	ps := make([]int, len(s), len(s))
	for i, _ := range ps {
		ps[i] = 1 // every letter is a palindrome of itself
	}

	var longestIdx int
	for i, _ := range s {
		for j := i; j < len(s); j++ {
			start := i + ((j - i) / 2)
			end := start + 1
			if i != j && (j-i)%2 == 0 {
				// odd-range length
				start -= 1
			}

			for ; 0 <= start && end < len(s); start, end = start-1, end+1 {
				if s[start] != s[end] {
					break
				}
				if ps[start] > (end-start)+1 {
					// previously seen palindrome that is longer that current palindrome, so short circuit
					break
				}
				ps[start] = (end - start) + 1
				if ps[start] > ps[longestIdx] {
					longestIdx = start
				}
			}
		}
	}
	return string(s[longestIdx : longestIdx+ps[longestIdx]])
}
