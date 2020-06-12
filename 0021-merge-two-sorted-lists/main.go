package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val <= l2.Val {
		head = &ListNode{l1.Val, nil}
		l1 = l1.Next
	} else {
		head = &ListNode{l2.Val, nil}
		l2 = l2.Next
	}
	last := head
	var v int
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil && l1.Val <= l2.Val {
			v = l1.Val
			l1 = l1.Next
		} else if l1 != nil && l2 != nil && l1.Val > l2.Val {
			v = l2.Val
			l2 = l2.Next
		} else if l1 != nil {
			v = l1.Val
			l1 = l1.Next
		} else if l2 != nil {
			v = l2.Val
			l2 = l2.Next
		} else {
			break
		}
		last.Next = &ListNode{v, nil}
		last = last.Next
	}
	return head
}
