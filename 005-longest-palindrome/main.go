package main

// indexes that make up a palindrome
type pair struct {
	start int
	end   int // inclusive
}

func (p pair) Len() int {
	return (p.end - p.start) + 1
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s
	}
	rs := []rune(s)
	ps := make(map[pair]bool)
	var longest pair
	for i, _ := range rs {
		for j := i + 1; j < len(rs); j++ {
			if _, ok := ps[pair{i, j}]; ok {
				// this is a previously seen palindrome so short-circuit
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
				if _, ok := ps[pair{start, end}]; ok {
					// previously seen palindrome, so short circuit
					break
				}
				ps[pair{start, end}] = true
				if end-start+1 > longest.Len() {
					longest = pair{start, end}
				}
			}
		}
	}
	return string(rs[longest.start : longest.end+1])
}
