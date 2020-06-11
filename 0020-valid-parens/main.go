package main

type stack []rune

func (s stack) push(r rune) stack {
	return append(s, r)
}

func (s stack) pop() stack {
	return s[:len(s)-1]
}

func (s stack) empty() bool {
	return len(s) == 0
}

func (s stack) peek() rune {
	return s[len(s)-1]
}

type pair struct {
	open  rune
	close rune
}

var pairs = []pair{
	{'(', ')'},
	{'[', ']'},
	{'{', '}'},
}

func isValid(s string) bool {
	var ss stack = make([]rune, 0, 0)
	for _, r := range s {
		for _, p := range pairs {
			if r == p.open {
				ss = ss.push(r)
				break
			} else if r == p.close && ss.peek() == p.open {
				ss = ss.pop()
				break
			} else if r == p.close && ss.peek() != p.open {
				return false
			}
		}
	}
	return ss.empty()
}
