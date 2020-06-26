package main

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{New(1, 2, 3, 4, 5), 2}, New(2, 1, 4, 3, 5)},
		{"Example 2", args{New(1, 2, 3, 4, 5), 3}, New(3, 2, 1, 4, 5)},
		{"k=len", args{New(1, 2, 3, 4, 5), 5}, New(5, 4, 3, 2, 1)},
		{"k=0", args{New(1, 2, 3, 4, 5), 0}, New(1, 2, 3, 4, 5)},
		{"k=1", args{New(1, 2, 3, 4, 5), 1}, New(1, 2, 3, 4, 5)},
		{"nil", args{nil, 3}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !Equals(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
