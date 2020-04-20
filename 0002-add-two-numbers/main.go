package main

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	s := "["
	for ; ln != nil; ln = ln.Next {
		s += strconv.Itoa(ln.Val)
		s += ","
	}
	s += "]"
	return s
}

func next(ln *ListNode) (int, *ListNode) {
	if ln == nil {
		return 0, nil
	}
	return ln.Val, ln.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	v1, l1 := next(l1)
	v2, l2 := next(l2)

	head := &ListNode{v1 + v2, nil}
	carry := 0 // carry forward a ten when necessary
	if head.Val >= 10 {
		head.Val -= 10
		carry = 1
	}
	last := head
	for l1 != nil || l2 != nil {
		v1, l1 = next(l1)
		v2, l2 = next(l2)
		sum := v1 + v2 + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}

		last.Next = &ListNode{sum, nil}
		last = last.Next
	}

	if carry == 1 {
		last.Next = &ListNode{1, nil}
	}
	return head
}
