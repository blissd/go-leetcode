package main

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{New(1, 2, 3, 4)}, New(2, 1, 4, 3)},
		{"Failure 1", args{New(1)}, New(1)},
		{"Failure 2", args{New(1, 2, 3)}, New(2, 1, 3)},
		{"Failure 3", args{New(1, 2, 3, 4, 5)}, New(2, 1, 4, 3, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !Equals(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
