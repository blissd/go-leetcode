package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s
	}
	rs := []rune(s)
	ps := make(map[int]int) // map from index to palindrome length
	var longestStart, longestEnd int
	for i, _ := range rs {
		for j := i + 1; j < len(rs); j++ {
			if length, ok := ps[i]; ok && length >= (j-i)+1 {
				// previously seen palindrome starts at this index
				continue
			}

			// range length is even
			start := i + ((j - i) / 2)
			end := start + 1
			if (j-i)%2 == 0 {
				// odd-range length
				start -= 1
			}

			// from middle of range, increment bounds finding all palindromes
			// hmm... should loop end when start==j and end==j?
			for ; 0 <= start && end < len(rs); start, end = start-1, end+1 {
				if rs[start] != rs[end] {
					break
				}
				if length, ok := ps[start]; ok && length >= (end-start)+1 {
					// previously seen palindrome, so short circuit
					break
				}
				ps[start] = (end - start) + 1
				if end-start+1 > (longestEnd-longestStart)+1 {
					longestStart, longestEnd = start, end
				}
			}
		}
	}
	return string(rs[longestStart : longestEnd+1])
}
