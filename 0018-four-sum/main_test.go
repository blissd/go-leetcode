package main

import (
	"reflect"
	"sort"
	"testing"
)

type resultList [][]int

func (t resultList) Len() int {
	return len(t)
}

func (t resultList) Less(i, j int) bool {
	for ii, jj := 0, 0; i < len(t[i]) && ii < len(t[i]) && jj < len(t[j]); ii, jj = ii+1, jj+1 {
		if t[i][ii] > t[j][jj] {
			return false
		}
	}
	return true
}

func (t resultList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{
			{-1, 0, 0, 1},
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.args.nums, tt.args.target)
			sort.Sort(resultList(got))
			want := tt.want
			sort.Sort(resultList(want))
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, want) {
				t.Errorf("fourSum() = %v, want %v", got, want)
			}
		})
	}
}
