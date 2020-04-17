package main

import (
	"strconv"
	"testing"
)
import "reflect"

func equals(t *testing.T, ln *ListNode, nums ...int) {
	listNums := []int{}
	for ; ln != nil; ln = ln.Next {
		listNums = append(listNums, ln.Val)
	}
	if !reflect.DeepEqual(ln, listNums) {
		t.Fatalf("expected %v, got %v", nums, listNums)
	}
}

func TestJoin(t *testing.T) {
	ln := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	sum := join(ln)
	if sum != 342 {
		t.Fatalf("expected 345, got %v\n", sum)
	}
}

func TestSplit1(t *testing.T) {
	equals(t, split(807), 7, 0, 8)

	var tests = []struct {
		n    int64
		nums []int
	}{
		{807, []int{7, 0, 8}},
		{10, []int{0, 1}},
		{1000000000000000000000000000001, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatInt(tt.n, 10), func(t *testing.T) {
			equals(t, split(tt.n), tt.nums...)
		})
	}
}
