/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast := head.Next.Next
	slow := head.Next
	flag := false
	for fast != nil && fast.Next != nil {
		if slow == fast {
			flag = true
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	if flag == false {
		return nil
	}
	for fast != head {
		fast = fast.Next
		head = head.Next
	}
	return fast
}