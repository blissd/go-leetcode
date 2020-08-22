package main

type Set map[string]struct{}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	results := make(Set)

	gen("()", results, n-1)

	parens := make([]string, len(results), len(results))
	i := 0
	for k := range results {
		parens[i] = k
		i++
	}

	return parens
}

func gen(parens string, results Set, n int) {
	if n == 0 {
		results[parens] = struct{}{}
		return
	}
	/// (())(())
	gen("()"+parens, results, n-1)
	gen("("+parens+")", results, n-1)
	gen(parens+"()", results, n-1)
}
