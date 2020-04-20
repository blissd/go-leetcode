package main

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"example 1a", args{[]int{1, 2, 3}, []int{}}, 2},
		{"example 1b", args{[]int{1, 2}, []int{3}}, 2},
		{"example 1c", args{[]int{1}, []int{2, 3}}, 2},
		{"example 1d", args{[]int{}, []int{1, 2, 3}}, 2},
		{"example 2a", args{[]int{1, 2, 3, 4}, []int{}}, 2.5},
		{"example 2b", args{[]int{1, 2, 3}, []int{4}}, 2.5},
		{"example 2c", args{[]int{1, 2}, []int{3, 4}}, 2.5},
		{"example 2d", args{[]int{1}, []int{2, 3, 4}}, 2.5},
		{"example 2e", args{[]int{}, []int{1, 2, 3, 4}}, 2.5},
		{"example 3", args{[]int{3, 4}, []int{1, 2, 5}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
