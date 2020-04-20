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
	for _, r := range p {
		switch {
		case r == '.':
			matchers = append(matchers, any{})
		case r == '*':
			// repeat previous match expression, not previous character
			matchers[len(matchers)-1] = repeat{matchers[len(matchers)-1]}
		default:
			matchers = append(matchers, exact(r))
		}
	}
	return matchers
}

func isMatch(s string, p string) bool {

	matchers := compile(p)
	var m int
	for _, r := range s {
		if m == len(matchers) {
			// we've used the last matcher but still have more characters to evaluate, so fail
			return false
		}
		for m < len(matchers) {
			result := matchers[m].match(r)
			if result == fail {
				return false
			}
			if result == unmatched {
				// no match, so advance to next matcher and try again
				m++
				if m == len(matchers) {
					return false
				}
				continue
			}
			if result == matched {
				m++
				break
			}
			if result == matched_repeat {
				break
			}
		}
	}

	// test matchers to see if they produce an error when they are unused
	if m < len(matchers) {
		for ; m < len(matchers); m++ {
			if matchers[m].unused() != nil {
				if matchers[m].match(rune(s[len(s)-1])) == matched {
					// a repeat of a character is followed by an exact match of a character,
					// then the exact match won't get evaluated.
					// So if there ureached matchers would match the last character, then don't fail.
					continue
				}
				return false
			}
		}
	}

	return true
}
