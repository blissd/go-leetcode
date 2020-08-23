package main

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	result := gen("(", 1, 0, n)
	return result
}

func gen(prefix string, open int, closed int, n int) []string {
	if closed == n {
		return []string{prefix}
	}

	var left []string
	if open < n {
		left = gen(prefix+"(", open+1, closed, n)
	}

	var right []string
	if closed < open {
		right = gen(prefix+")", open, closed+1, n)
	}

	return append(left, right...)
}
