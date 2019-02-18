/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	fast := head.Next.Next
	slow := head.Next

	for fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return false

}