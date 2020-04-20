package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	rs := []rune(s)
	var chars = make(map[rune]int)
	start := 0

	longest := 0
	for end, c := range rs {
		if dupeIdx, ok := chars[c]; ok {
			// duplicate found, so fast-forward sequence start to first occurrence of dupe.
			// Remove seen characters on the way.
			for ; start <= dupeIdx; start++ {
				delete(chars, rs[start])
			}
		}
		chars[c] = end
		longest = max(longest, (end-start)+1)
	}
	return longest
}
