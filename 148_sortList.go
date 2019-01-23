/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, mid, fast := head, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		mid = slow
		slow = slow.Next
	}

	if fast != nil {
		mid = slow
		slow = slow.Next
	}
	mid.Next = nil

	list1 := sortList(head)
	list2 := sortList(slow)

	var newHead ListNode
	tmp := &newHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			tmp = tmp.Next
			list1 = list1.Next
		} else {
			tmp.Next = list2
			tmp = tmp.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		tmp.Next = list1
	}

	if list2 != nil {
		tmp.Next = list2
	}
	return newHead.Next
}