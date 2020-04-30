package main

import (
	"reflect"
	"sort"
	"testing"
)

type threeList [][]int

func (t threeList) Len() int {
	return len(t)
}

func (t threeList) Less(i, j int) bool {
	if t[i][0] < t[j][0] {
		return true
	}
	if t[i][0] == t[j][0] && t[i][1] < t[j][1] {
		return true
	}
	if t[i][0] == t[j][0] && t[i][1] == t[j][1] && t[i][2] < t[j][2] {
		return true
	}
	return false
}

func (t threeList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"example 1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{"failure 1", args{[]int{0, 0}}, [][]int{}},
		{"failure 2", args{[]int{1, 2, -2, -1}}, [][]int{}},
		{"failure 3", args{[]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}}, [][]int{{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}}},
	}

	// sort wanted results
	for _, tt := range tests {
		sort.Sort(threeList(tt.want))
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)
			sort.Sort(threeList(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
