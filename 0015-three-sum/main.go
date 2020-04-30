package main

import (
	"sort"
)

type three struct {
	a, b, c int
}

func (t three) sort() three {
	tt := []int{t.a, t.b, t.c}
	sort.Ints(tt)
	return three{tt[0], tt[1], tt[2]}
}

func threeSum(nums []int) [][]int {
	threes := make(map[three]bool) // a unique set
	for _, n := range nums {
		for _, t := range findThrees(nums, n) {
			threes[t] = true
		}
	}

	result := make([][]int, 0, len(threes))
	for k, _ := range threes {
		result = append(result, []int{k.a, k.b, k.c})
	}
	return result
}

func findThrees(nums []int, target int) []three {
	// map from value to index
	threes := []three{}

	counts := count(nums)
	for n, ncount := range counts {
		nn := (-target) - n

		if (n == target || nn == target || n == nn) && ncount < 2 {
			continue
		}

		if n == target && nn == target && ncount < 3 {
			continue
		}
		if nncount, ok := counts[nn]; ok && nncount > 0 {
			threes = append(threes, three{target, n, nn}.sort())
		}
	}

	return threes
}

func count(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, n := range nums {
		c, _ := counts[n]
		counts[n] = c + 1
	}
	return counts
}
