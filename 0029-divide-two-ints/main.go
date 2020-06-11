package main

import "math"

func divide(dividend int, divisor int) int {
	d := dividend
	if d < 0 {
		d = -d
	}

	r := divisor
	if r < 0 {
		r = -r
	}

	i, v := 0, 0
	for ; v < d; i, v = i+1, v+r {
		if v+r > d {
			break
		}
	}
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		i = -i
	}

	if i < math.MinInt32 || math.MaxInt32 < i {
		return math.MaxInt32
	}

	return i
}
