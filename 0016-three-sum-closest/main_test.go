package main

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{[]int{-1, 2, 1, -4}, 1}, 2},
		{"failure 1", args{[]int{1, 1, 1, 0}, 100}, 3},
		{"failure 2", args{[]int{0, 1, 2}, 0}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
