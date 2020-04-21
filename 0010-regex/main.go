package main

type result int

const (
	abort           result = 1
	advance_matcher        = 2
	advance_letter         = 4
)

// a rule for matching a rune
type matcher interface {
	// is this rune matched? Nil returned if matched.
	match(c uint8) result
}

// Matches a specific character.
type once uint8

// Matches any character.
type any struct{}

type composite struct {
	m matcher
}

// Matches another matcher zero or more times
type repeat composite

// match any character except one
type invert composite

// match exactly
type exact composite

func (m exact) match(c uint8) result {
	return m.m.match(c)
}

func (m once) match(c uint8) result {
	if c == uint8(m) {
		return advance_letter | advance_matcher
	}
	return abort
}

func (m any) match(_ uint8) result {
	return advance_letter | advance_matcher
}

func (m invert) match(c uint8) result {
	r := m.m.match(c)
	if r&advance_letter == advance_letter {
		return advance_matcher
	}
	if r&advance_matcher == advance_matcher {
		return advance_letter
	}
	return abort
}

func (m repeat) match(c uint8) result {
	r := m.m.match(c)
	if r == abort {
		return advance_matcher
	}
	return advance_letter
}

func compile(p string) []matcher {
	matchers := make([]matcher, 0, len(p))
	for _, r := range p {
		switch {
		case r == '.':
			matchers = append(matchers, exact{any{}})
		case r == '*':
			m := repeat{matchers[len(matchers)-1]}
			// Consolidate adjacent repeating matchers of the same character
			if len(matchers) > 2 && matchers[len(matchers)-2] == m {
				continue
			}
			// Repeat previous match expression, not previous character
			matchers[len(matchers)-1] = m
		default:
			if len(matchers) > 1 && matchers[len(matchers)-1] == (repeat{exact{any{}}}) {
				matchers[len(matchers)-1] = invert{repeat{once(r)}}
			}
			//if len(matchers) > 1 {
			//switch matchers[len(matchers)-1].(type) {
			//case any:
			//	matchers[len(matchers)-1] = not{exact(r)}
			//}
			//}
			matchers = append(matchers, exact{once(r)})
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

	var exactCount int // count of exact matches in pattern
	for m := 0; m < len(matchers); m++ {
		if _, ok := matchers[m].(exact); ok {
			exactCount++
		}
	}
	var i, m int
	for m < len(matchers) && i < len(s) {
		if _, ok := matchers[m].(exact); ok {
			exactCount--
		} else if i+exactCount == len(s) {
			// we are on a repeating matcher but have unused exact matchers.
			// Advance to next exact matcher and start using that.
			for ; m < len(matchers); m++ {
				if _, ok := matchers[m].(exact); ok {
					break
				}
			}
		}
		mm := matchers[m]
		result := mm.match(s[i])
		if result == abort {
			return false
		}
		if result&advance_matcher == advance_matcher {
			m++
		}
		if result&advance_letter == advance_letter {
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
