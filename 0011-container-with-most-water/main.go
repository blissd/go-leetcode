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

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var area int
	for left != right {
		area = max(area, (right-left)*min(height[left], height[right]))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return area
}
