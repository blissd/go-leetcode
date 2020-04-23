package main

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"failure 1", args{[]int{1, 2}}, 1},
		{"random 1", args{[]int{1, 2, 3}}, 2},
		{"random 2", args{[]int{27, 46, 90, 70, 48, 39, 61, 4, 55, 75, 67, 72, 30, 40, 25, 2, 6, 89, 8, 1}}, 1335},
		{"random 3", args{[]int{10, 78, 52, 2, 5, 48, 49, 25}}, 245},
		{"random 4", args{[]int{79, 88, 99, 30, 2, 99, 55, 15}}, 395},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
