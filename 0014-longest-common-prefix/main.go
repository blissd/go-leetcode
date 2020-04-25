package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var i int
	for ; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:i]
}
