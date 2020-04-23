package main

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxArea(height []int) int {
	var maxArea int
	for x1, h1 := range height {
		for x2, h2 := range height {
			if x1 == x2 {
				continue
			}

			width := abs(x2 - x1)
			height := min(h1, h2)
			maxArea = max(maxArea, width*height)
		}
	}
	return maxArea
}
