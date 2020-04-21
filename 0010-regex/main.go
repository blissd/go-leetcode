package main

import (
	"fmt"
)

type result int

const (
	// matching has failed
	fail result = iota + 1
	// No match, but instead of failing advance to next matcher and evaluate again
	unmatched
	// Match, but don't advance to next matcher.
	matched_repeat
	// Match. Advance to next matcher, but don't evalute current rune again
	matched
)

// a rule for matching a rune
type matcher interface {
	// is this rune matched? Nil returned if matched.
	match(c rune) result
	// error for when matcher isn't evaluated against any runes. Nil for no error.
	unused() error
}

// Matches a specific rune.
type exact rune

// Matches any rune.
type any struct{}

// Matches another matcher zero or more times
type repeat struct {
	m matcher
}

func (m exact) match(c rune) result {
	if c == rune(m) {
		return matched
	}
	return fail
}

func (m exact) unused() error {
	return fmt.Errorf("unmatched rune: %v", rune(m))
}

func (m any) match(_ rune) result {
	return matched
}

func (m any) unused() error {
	return nil
}

func (m repeat) match(c rune) result {
	if m.m.match(c) != fail {
		return matched_repeat
	}
	return unmatched
}

func (m repeat) unused() error {
	return nil
}

func compile(p string) []matcher {
	matchers := make([]matcher, 0, len(p))
	for ii, r := range p {
		switch {
		case r == '.':
			matchers = append(matchers, any{})
		case r == '*':
			// Consolidate adjacent repeating matchers of the same character
			if matchers[len(matchers)-1].match(rune(p[ii-1])) == matched_repeat {
				continue
			}
			// Repeat previous match expression, not previous character
			matchers[len(matchers)-1] = repeat{matchers[len(matchers)-1]}
		default:
			matchers = append(matchers, exact(r))
		}
	}

	// re-order matchers so an exact(x) and repeat(exact(x)) are adjacent, then the exact comes first
	for i := 1; i < len(matchers); i++ {
		if matchers[i-1] == (repeat{matchers[i]}) {
			matchers[i-1], matchers[i] = matchers[i], matchers[i-1]
			i = 0 // restart sort from beginning
		}
	}

	return matchers
}

func isMatch(s string, p string) bool {

	matchers := compile(p)
	var m int
	for i, r := range s {
		i = i
		if m == len(matchers) {
			// we've used the last matcher but still have more characters to evaluate, so fail
			return false
		}

		result := matchers[m].match(r)

		if result == fail {
			return false
		} else if result == matched {
			m++
		} else if result == unmatched {
			// no match, so advance matchers until match or failure
			for m += 1; m < len(matchers); m++ {
				result := matchers[m].match(r)
				if result == fail {
					return false
				}
				if result == unmatched {
					continue
				}
				if result == matched_repeat {
					break
				}
				if result == matched {
					m++
					break
				}
			}
		} else if result == matched_repeat {
			// must match shortest set of repeating characters
			for mm := m + 1; mm < len(matchers); mm++ {
				if matchers[mm].match(r) != matched_repeat {
					break
				}
				m = mm
			}
		}
	}

	// test matchers to see if they produce an error when they are unused
	if m < len(matchers) {
		for ; m < len(matchers); m++ {
			if matchers[m].unused() != nil {
				return false
			}
		}
	}

	return true
}
