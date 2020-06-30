package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	// build mapping from value to indexes in nums that have that value
	indexes := make(map[int][]int, len(nums))
	for i, v := range nums {
		if ii, ok := indexes[v]; ok {
			indexes[v] = append(ii, i)
		} else {
			indexes[v] = []int{i}
		}
	}

	results := make(map[[4]int]struct{}, 10)
	exists := struct{}{}

	for i, j, k := 0, 1, 2; i+3 < len(nums); j++ {
		if j == k {
			j = i + 1
			k = k + 1
		}
		if k == len(nums) {
			i = i + 1
			j = i + 1
			k = j + 1
		}
		sum := nums[i] + nums[j] + nums[k]
		key := target - sum
		if ll, ok := indexes[key]; ok {
			// we have a match, but must make sure we have an unused index
			for _, v := range ll {
				if v != i && v != j && v != k {
					result := [4]int{nums[i], nums[j], nums[k], key}
					sort.Ints(result[:])
					results[result] = exists
					break
				}
			}
		}
	}

	keys := make([][]int, 0, len(results))
	for k := range results {
		keys = append(keys, []int{k[0], k[1], k[2], k[3]})
	}

	return keys
}
