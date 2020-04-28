package main

import (
	"errors"
	"sort"
)

var ErrNoSum = errors.New("no sum")

type three struct {
	a, b, c int
}

func threeSum(nums []int) [][]int {
	rs := make(map[three]bool)
	for _, n := range nums {
		if a, b, err := twoSum(nums, -1); err == nil {
			r := []int{a, b, n}
			sort.Ints(r)
			rs[three{r[0], r[1], r[2]}] = true
		}
	}

	keys := make([][]int, 0, len(rs))
	for k, _ := range rs {
		keys = append(keys, []int{k.a, k.b, k.c})
	}
	return keys
}

func twoSum(nums []int, target int) (int, int, error) {
	// map from value to index
	lookup := count(nums)
	for i, n := range nums {
		// compute value we want to lookup index for
		nn := target - n
		if ii, ok := lookup[nn]; ok && ii != i {
			return nums[ii], nums[i], nil
		}

		lookup[n] = i
	}

	return 0, 0, ErrNoSum
}

func count(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, n := range nums {
		c, _ := counts[n]
		counts[n] = c + 1
	}
	return counts
}
