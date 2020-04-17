package main

func median(nums []int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	} else if len(nums)%2 != 0 {
		return float64(nums[len(nums)/2])
	} else {
		// compute mean of two mids
		mid := int(len(nums) / 2)
		lo, hi := float64(nums[mid]), float64(nums[mid+1])
		return (lo + hi) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums = make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))
	i1 := 0
	i2 := 0
	for i, _ := range nums {
		if i1 >= len(nums1) {
			nums[i] = nums2[i2]
			i2++
		} else if i2 >= len(nums2) {
			nums[i] = nums1[i1]
			i1++
		} else if nums1[i1] < nums2[i2] {
			nums[i] = nums1[i1]
			i1++
		} else {
			nums[i] = nums2[i2]
			i2++
		}
	}
	return median(nums)
}
