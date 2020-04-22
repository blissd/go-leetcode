package main

// what to do after evaluating a matcher
type result int

const (
	advance_matcher = 2
	advance_letter  = 4
)

// a rule for matching a character
type matcher interface {
	match(c uint8) result
}

// match one time
type single uint8

// Matches another matcher zero or more times
type repeat struct {
	m matcher
}

func (m single) match(c uint8) result {
	if uint8(m) == '.' || c == uint8(m) {
		return advance_letter | advance_matcher
	}
	return advance_matcher
}

func (m repeat) match(c uint8) result {
	r := m.m.match(c)
	if r == advance_matcher {
		return advance_matcher
	}
	return advance_letter
}

func compile(p string) []matcher {
	matchers := make([]matcher, 0, len(p))
	for _, r := range p {
		if r == '*' {
			matchers[len(matchers)-1] = repeat{matchers[len(matchers)-1]}
		} else {
			matchers = append(matchers, single(r))
		}
	}
	return matchers
}

func isMatch2(s string, matchers []matcher) bool {

	//if len(s) == 0 || len(matchers) == 0 {
	//	return false
	//}

	var i, m int
	for m < len(matchers) {
		if m == len(matchers)-1 && i == len(s) {
			return true
		}
		if i == len(s) {
			return false
		}
		mm := matchers[m]
		result := mm.match(s[i])
		if result == advance_matcher|advance_letter {
			// exact single letter match
			m++
			i++
		} else if result == advance_letter {
			// Letter is advancing, but matcher is not.
			// This means we are evaluating a repeat matcher so we must see if
			// the subsequent matchers can fully match the string from this position
			if m+1 < len(matchers) && isMatch2(s[i:], matchers[m+1:]) {
				return true
			}
			i++
		} else if result == advance_matcher {
			m++
		}
	}

	return m == len(matchers) && i == len(s)
}

func isMatch(s string, p string) bool {
	return isMatch2(s, compile(p))
}
