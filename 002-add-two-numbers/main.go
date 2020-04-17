package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func join(ln *ListNode) int64 {
	shift := int64(1)
	var sum int64
	for n := ln; n != nil; n = n.Next {
		sum += int64(n.Val) * shift
		shift *= 10
	}
	return sum
}

func split(num int64) *ListNode {
	last := &ListNode{int(num - ((num / 10) * 10)), nil}
	head := last
	num /= 10
	for num > 0 {
		digit := int(num - ((num / 10) * 10))
		last.Next = &ListNode{digit, nil}
		last = last.Next
		num /= 10
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return split(join(l1) + join(l2))
}

func main() {
	join(&ListNode{9, nil})
	split(18)
	//ln := &ListNode{1, &ListNode{4, &ListNode{3, nil}}}
	//ln := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	//fmt.Println(join(ln))
	//
	//fmt.Println()
}
