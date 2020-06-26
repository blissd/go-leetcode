package main

import (
	"fmt"
	"reflect"
	"strings"
)

// structure used in many exercises
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	var s strings.Builder
	for n != nil {
		s.WriteString(fmt.Sprintf("%v ", n.Val))
		n = n.Next
	}
	return s.String()
}

func New(ns ...int) *ListNode {
	if len(ns) == 0 {
		return nil
	}

	head := &ListNode{ns[0], nil}
	last := head
	for _, n := range ns[1:] {
		last.Next = &ListNode{n, nil}
		last = last.Next
	}
	return head
}

func Equals(l1 *ListNode, l2 *ListNode) bool {
	var l1nums []int
	for ; l1 != nil; l1 = l1.Next {
		l1nums = append(l1nums, l1.Val)
	}
	var l2nums []int
	for ; l2 != nil; l2 = l2.Next {
		l2nums = append(l2nums, l2.Val)
	}
	return reflect.DeepEqual(l1nums, l2nums)
}
