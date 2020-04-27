package main

import "sort"

type three struct {
	a, b, c int
}

func threeSum(nums []int) [][]int {
	rs := make(map[three]bool)
	for _, n := range nums {
		r := twoSum(nums, -n)
		if r != nil {
			r = append(r, n)
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

func twoSum(nums []int, target int) []int {
	// map from value to index
	lookup := make(map[int]int)
	for i, n := range nums {
		// compute value we want to lookup index for
		nn := target - n
		if ii, ok := lookup[nn]; ok && ii != i {
			return []int{nums[ii], nums[i]} // lower index first
		}

		lookup[n] = i
	}

	return nil
}
