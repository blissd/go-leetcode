package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	first := &ListNode{head.Val, nil}
	last := first
	buf := make([]*ListNode, n+1, n+1)
	buf[0] = first
	i := 1
	for node := head.Next; node != nil; i, node = i+1, node.Next {
		next := &ListNode{node.Val, nil}
		last.Next = next
		last = next
		j := i % len(buf)
		buf[j] = last
	}
	if n > i {
		return nil
	}
	if n == i {
		return first.Next
	}
	removeIdx := (i - n) % len(buf)
	beforeIdx := (i - n - 1) % len(buf)
	buf[beforeIdx].Next = buf[removeIdx].Next
	return first
}
