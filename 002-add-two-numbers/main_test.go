package main

import (
	"fmt"
	"testing"
)
import "reflect"

func equals(t *testing.T, ln *ListNode, nums []int) {
	var listNums []int
	for ; ln != nil; ln = ln.Next {
		listNums = append(listNums, ln.Val)
	}
	if !reflect.DeepEqual(nums, listNums) {
		t.Fatalf("expected %v, got %v", nums, listNums)
	}
}

func newListNode(ns []int) *ListNode {
	var head, last *ListNode
	for _, n := range ns {
		if head == nil {
			head = &ListNode{n, nil}
			last = head
		} else {
			last.Next = &ListNode{n, nil}
			last = last.Next
		}
	}
	return head
}

func add(t *testing.T, ns1 []int, ns2 []int, sum []int) {
	equals(t, addTwoNumbers(newListNode(ns1), newListNode(ns2)), sum)
}

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		ns1 []int
		ns2 []int
		sum []int
	}{
		{[]int{5}, []int{5}, []int{0, 1}},
		{[]int{7, 0, 8}, []int{0}, []int{7, 0, 8}},
		{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, []int{5, 6, 4}, []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for ii, tt := range tests {
		t.Run(fmt.Sprintf("test %v", ii), func(t *testing.T) {
			add(t, tt.ns1, tt.ns2, tt.sum)
		})
	}
}
