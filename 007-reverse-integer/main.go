package main

import "math"

func reverse(x int) int {
	var reversed int
	for ; x != 0; x /= 10 {
		reversed *= 10
		reversed += x - ((x / 10) * 10)
	}
	if math.MaxInt32 < reversed || reversed < math.MinInt32 {
		return 0
	}
	return reversed
}
