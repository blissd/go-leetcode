package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	first := &ListNode{head.Next.Val, &ListNode{head.Val, nil}}
	last := first.Next
	head = head.Next.Next

	for head != nil {
		if head.Next != nil {
			last.Next = &ListNode{head.Next.Val, &ListNode{head.Val, nil}}
			last = last.Next.Next
			head = head.Next.Next
		} else {
			last.Next = &ListNode{head.Val, nil}
			last = last.Next
			head = head.Next
		}
	}
	return first
}
