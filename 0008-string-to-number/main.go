package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	// start index of number in string, including leading + or - sign.
	var start int
	for i, c := range s {
		if unicode.IsSpace(c) {
			continue
		}
		if unicode.IsNumber(c) || c == '-' || c == '+' {
			start = i
			break
		}
		return 0
	}
	neg := s[start] == '-'
	if s[start] == '-' || s[start] == '+' {
		start += 1
	}

	// now we know where the number starts, keep converting digits until the first non-numeric character
	var sum int
	for _, c := range s[start:] {
		if !unicode.IsNumber(c) {
			break
		}
		sum *= 10
		sum += int(c - '0')

		if neg && sum > math.MaxInt32+1 {
			return math.MinInt32
		} else if !neg && sum > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if neg {
		return -sum
	}
	return sum
}
