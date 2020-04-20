package main

func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	// each element represents the length of the palindrome at that index in the string
	ps := make([]int, len(s), len(s))

	var longestIdx int
	for i, _ := range s {
		for j := i + ps[i]; j < len(s); j++ {
			start := i + ((j - i) / 2)
			end := start + 1
			if i != j && (j-i)%2 == 0 {
				// odd-range length
				start -= 1
			}

			for ; 0 <= start && end < len(s); start, end = start-1, end+1 {
				// not a palindrome or
				// previously seen palindrome that is longer that current palindrome, so short circuit
				if s[start] != s[end] || ps[start] > end-start {
					break
				}
				ps[start] = end - start
				if ps[start] > ps[longestIdx] {
					longestIdx = start
				}
			}
		}
	}
	if ps[longestIdx] == 0 {
		return s[:1] // if longest length is zero, then all letters are unique
	}
	return s[longestIdx : longestIdx+ps[longestIdx]+1]
}
