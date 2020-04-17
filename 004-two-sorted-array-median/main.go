package main

// iterate nums1 and nums2 in ascending order
func next(nums1 []int, i1 int, nums2 []int, i2 int) (v int, i1next int, i2next int) {
	if i1 == len(nums1) {
		return nums2[i2], i1, i2 + 1
	} else if i2 == len(nums2) {
		return nums1[i1], i1 + 1, i2
	} else if nums1[i1] <= nums2[i2] {
		return nums1[i1], i1 + 1, i2
	} else {
		return nums2[i2], i1, i2 + 1
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	mid := length / 2
	if length%2 != 0 {
		mid += 1
	}
	var v1, i1, i2 int
	for i := 0; i < mid; i++ {
		v1, i1, i2 = next(nums1, i1, nums2, i2)
	}

	v2 := v1
	if length%2 == 0 {
		v2, _, _ = next(nums1, i1, nums2, i2)
	}
	return (float64(v1) + float64(v2)) / 2
}
