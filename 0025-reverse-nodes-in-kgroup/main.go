package main

func atLeast(head *ListNode, n int) bool {
	for i := 0; i < n; head, i = head.Next, i+1 {
		if head == nil {
			return false
		}
	}
	return true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if !atLeast(head, k) {
		return head
	}

	var first *ListNode
	last := first

	for atLeast(head, k) {
		// build reversed segment
		end := &ListNode{head.Val, nil}
		start := end
		head = head.Next
		for i := 0; i < k-1; i++ {
			start = &ListNode{head.Val, start}
			head = head.Next
		}
		if last == nil {
			first = start
			last = end
		} else {
			last.Next = start
			last = end
		}
	}

	last.Next = head
	return first
}
