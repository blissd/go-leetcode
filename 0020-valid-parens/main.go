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

func isValid(s string) bool {
	var ss stack = make([]rune, 0, 0)
	for _, r := range s {
		if r == '{' || r == '(' || r == '[' {
			ss = ss.push(r)
		} else if r == '}' || r == ')' || r == ']' {
			if ss.empty() {
				return false
			}
			if r == '}' && ss.peek() != '{' {
				return false
			}
			if r == ']' && ss.peek() != '[' {
				return false
			}
			if r == ')' && ss.peek() != '(' {
				return false
			}
			ss = ss.pop()
		}
	}
	return ss.empty()
}
