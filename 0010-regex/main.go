package main

// What to do after evaluating a matcher
type result int

// Result of a evaluating a matcher and a character.
//Either the matcher will be advanced, the letter, or both.
const (
	abort           result = 1
	advance_matcher result = 2
	advance_letter  result = 4
)

// A rule for matching a character
type matcher interface {
	match(c uint8) result
}

// Match one time
type single uint8

func (m single) match(c uint8) result {
	if uint8(m) == '.' || c == uint8(m) {
		return advance_letter | advance_matcher
	}
	return abort
}

// Matches zero or more times
type repeat uint8

func (m repeat) match(c uint8) result {
	if uint8(m) == '.' || c == uint8(m) {
		return advance_letter
	}
	return advance_matcher
}

func compile(p string) []matcher {
	matchers := make([]matcher, 0, len(p))
	for i, r := range p {
		if r == '*' {
			matchers[len(matchers)-1] = repeat(p[i-1])
		} else {
			matchers = append(matchers, single(r))
		}
	}
	return matchers
}

type matchKey struct {
	i int // string index
	m int // matcher index
}

func isMatch2(s string, i int, matchers []matcher, m int, cache map[matchKey]bool) bool {
	for m < len(matchers) && i < len(s) {
		if match, ok := cache[matchKey{i: i, m: m}]; ok {
			return match
		}

		mm := matchers[m]
		rs := mm.match(s[i])
		if rs == abort {
			return false
		}
		if rs == advance_matcher|advance_letter {
			// exact single letter match
			m++
			i++
		} else if rs == advance_letter {
			// Letter is advancing, but matcher is not.
			// This means we are evaluating a repeat matcher so we must see if
			// the subsequent matchers can fully match the string from this position
			//if match, ok := cache[matchKey{i:i, m:m+1}]; ok {
			//	return match
			//}

			if m+1 < len(matchers) && isMatch2(s, i, matchers, m+1, cache) {
				cache[matchKey{i: i, m: m + 1}] = true
				return true
			} else {
				cache[matchKey{i: i, m: m + 1}] = false
			}
			i++
		} else if rs == advance_matcher {
			m++
		}
	}
	// if all subsequent matchers are repeats, then match is successful
	if i == len(s) {
		for ; m < len(matchers); m++ {
			if _, ok := matchers[m].(single); ok {
				return false
			}
		}
		return true
	}
	return false
}

func isMatch(s string, p string) bool {
	return isMatch2(s, 0, compile(p), 0, make(map[matchKey]bool))
}
