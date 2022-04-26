package main

import "math"

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	// steps[i] is the lowest cumulative cost to use the i'th step.
	steps := make([]int, len(cost))
	for i, _ := range steps {
		steps[i] = math.MaxInt
	}
	steps[0], steps[1] = cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		steps[i] = cost[i] + min(steps[i-1], steps[i-2])
	}

	return min(steps[len(steps)-1], steps[len(steps)-2])
}

// oh, golang!
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
