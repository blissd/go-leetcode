package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	lowest := -1

	for i := range lists {
		if lowest == -1 && lists[i] != nil {
			lowest = i
		} else if lowest != -1 && lists[i] != nil && lists[i].Val < lists[lowest].Val {
			lowest = i
		}
	}

	if lowest == -1 {
		return nil
	}

	head := &ListNode{lists[lowest].Val, nil}
	last := head
	lists[lowest] = lists[lowest].Next

	for {
		lowest = -1
		for i := range lists {
			if lowest == -1 && lists[i] != nil {
				lowest = i
			} else if lowest != -1 && lists[i] != nil && lists[i].Val < lists[lowest].Val {
				lowest = i
			}
		}
		if lowest == -1 {
			break
		}

		last.Next = &ListNode{lists[lowest].Val, nil}
		last = last.Next
		lists[lowest] = lists[lowest].Next
	}

	return head
}
