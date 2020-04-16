package main

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	// map from value to index
	lookup := make(map[int]int)
	for i, n := range nums {
		// compute value we want to lookup index for
		nn := target - n
		if ii, ok := lookup[nn]; ok && ii != i {
			return []int{ii, i} // lower index first
		}

		lookup[n] = i
	}

	return nil
}
