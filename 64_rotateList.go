/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	len := 1
	iterator := head
	for iterator.Next != nil {
		len++
		iterator = iterator.Next
	}
	k = k % len
	if k == 0 {
		return head
	}
	k = len - k

	iterator.Next = head

	for k > 1 {
		head = head.Next
		k--
	}
	iterator = head.Next
	head.Next = nil
	return iterator
}