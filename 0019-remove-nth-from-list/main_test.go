package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func list(ii ...int) *ListNode {
	var head *ListNode
	for i := len(ii) - 1; i >= 0; i-- {
		head = &ListNode{ii[i], head}
	}
	return head
}

func (n *ListNode) String() string {
	var s strings.Builder
	for n != nil {
		s.WriteString(fmt.Sprintf("%v ", n.Val))
		n = n.Next
	}
	return s.String()
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"example 1", args{list(1, 2, 3, 4, 5), 2}, list(1, 2, 3, 5)},
		{"example 1a", args{list(1, 2, 3, 4, 5), 1}, list(1, 2, 3, 4)},
		{"example 1b", args{list(1, 2, 3, 4, 5), 3}, list(1, 2, 4, 5)},
		{"example 1c", args{list(1, 2, 3, 4, 5), 4}, list(1, 3, 4, 5)},
		{"example 1c", args{list(1, 2, 3, 4, 5), 5}, list(2, 3, 4, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
