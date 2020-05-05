package main

import (
	"sort"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := -target
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		start := i + 1
		end := len(nums) - 1
		for start != end {
			b := nums[start]
			c := nums[end]
			sum := a + b + c
			if target == sum {
				return sum
			}
			if abs(target-sum) < abs(target-closest) {
				closest = sum
			}
			if target-sum > 0 {
				end--
			} else {
				start++
			}
		}
	}

	return closest
}
