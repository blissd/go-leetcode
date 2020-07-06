package main

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{3, 2, 2, 3}, 3}, []int{2, 2}},
		{"Example 2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, []int{0, 1, 3, 0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if !reflect.DeepEqual(tt.args.nums[:got], tt.want) || got != len(tt.want) {
				t.Errorf("removeElement() = %v, want %v", tt.args.nums[:got], tt.want)
			}
		})
	}
}
