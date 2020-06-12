package main

import (
	"reflect"
	"testing"
)

func equals(t *testing.T, l1 *ListNode, l2 *ListNode) {
	var l1nums []int
	for ; l1 != nil; l1 = l1.Next {
		l1nums = append(l1nums, l1.Val)
	}
	var l2nums []int
	for ; l2 != nil; l2 = l2.Next {
		l2nums = append(l2nums, l2.Val)
	}
	if !reflect.DeepEqual(l1nums, l2nums) {
		t.Errorf("expected %v, got %v", l2nums, l1nums)
	}
}

func nn(first int, ns ...int) *ListNode {
	head := &ListNode{first, nil}
	last := head
	for _, n := range ns {
		last.Next = &ListNode{n, nil}
		last = last.Next
	}
	return head
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"example 1", args{nn(1, 2, 4), nn(1, 3, 4)}, nn(1, 1, 2, 3, 4, 4)},
		{"failure 1", args{nn(2), nn(1)}, nn(1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equals(t, mergeTwoLists(tt.args.l1, tt.args.l2), tt.want)
		})
	}
}
