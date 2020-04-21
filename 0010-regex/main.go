package main

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
	match(c uint8) result
}

// Matches a specific character.
type exact uint8

// Matches any character.
type any struct{}

type composite struct {
	m matcher
}

// Matches another matcher zero or more times
type repeat composite

// match any character except one
type except composite

func (m exact) match(c uint8) result {
	if c == uint8(m) {
		return matched
	}
	return fail
}

func (m any) match(_ uint8) result {
	return matched
}

func (m except) match(c uint8) result {
	if m.m.match(c) == matched_repeat || m.m.match(c) == matched {
		return unmatched
	}
	return matched
}

func (m repeat) match(c uint8) result {
	if m.m.match(c) != fail {
		return matched_repeat
	}
	return unmatched
}

func compile(p string) []matcher {
	matchers := make([]matcher, 0, len(p))
	for _, r := range p {
		switch {
		case r == '.':
			matchers = append(matchers, any{})
		case r == '*':
			m := repeat{matchers[len(matchers)-1]}
			// Consolidate adjacent repeating matchers of the same character
			if len(matchers) > 2 && matchers[len(matchers)-2] == m {
				continue
			}
			// Repeat previous match expression, not previous character
			matchers[len(matchers)-1] = m
		default:
			if len(matchers) > 1 {
				switch matchers[len(matchers)-1].(type) {
				case any:
					matchers[len(matchers)-1] = except{exact(r)}
				}
			}
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
	var i, m int
	for m < len(matchers) && i < len(s) {
		//if i == len(s) {
		//	return false // unused matchers remaining
		//}
		mm := matchers[m]
		result := mm.match(s[i])
		if result == fail {
			return false
		} else if result == unmatched {
			m++
			continue
		} else if result == matched {
			m++
			i++
		} else if result == matched_repeat {
			// advance matcher
			for j := m + 1; j < len(matchers); j++ {
				rs := matchers[j].match(s[i])
				if rs == matched_repeat || rs == unmatched {
					m = j
					continue
				}
				//else if matchers[j].match(s[i]) == matched {
				//	m = j  // +1 because this matcher is satisfied by s[i]
				//	break
				//}
				break
			}
			i++
		}
	}

	for ; m < len(matchers); m++ {
		if _, ok := matchers[m].(exact); ok {
			return false
		}
	}
	return i == len(s)
}
