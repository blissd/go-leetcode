package main

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	// cumulative cost of two previous steps
	downOne, downTwo := cost[0], cost[1]
	step := 0

	for i := 2; i < len(cost); i++ {
		step = cost[i] + min(downOne, downTwo)
		downOne, downTwo = downTwo, step
	}

	return min(downOne, downTwo)
}

// oh, golang!
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
