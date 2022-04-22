package coinchange

func coinChange(coins []int, amount int) int {
	counts := make([]int, amount+1)
	for i := range counts {
		counts[i] = amount + 1
	}

	counts[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			counts[i] = min(counts[i], counts[i-coin]+1)
		}
	}

	if counts[amount] == amount+1 {
		return -1
	}
	return counts[amount]
}

// oh, golang!
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
