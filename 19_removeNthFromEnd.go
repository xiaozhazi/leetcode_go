package main

import (
	"fmt"
)

type ListNode struct {
	Var  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	fast := head
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}
