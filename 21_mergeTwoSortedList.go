package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var newHead ListNode
	tmp := &newHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next, tmp = l1, l1
			l1 = l1.Next
		} else {
			tmp.Next, tmp = l2, l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return newHead.Next
}