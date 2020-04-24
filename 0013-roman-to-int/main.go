package main

import "strings"

// represents a lower level and the numeral for that level
type band struct {
	i int    // integer
	n string // numeral
}

var bands = []band{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToInt(s string) int {
	var num int
	for _, b := range bands {
		for strings.HasPrefix(s, b.n) {
			num += b.i
			s = s[len(b.n):]
		}
	}
	return num
}
