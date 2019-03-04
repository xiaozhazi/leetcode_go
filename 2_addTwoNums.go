/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	add := 0
	head := &ListNode{
		Val:  (l1.Val + l2.Val) % 10,
		Next: nil,
	}
	ans := head
	add = (l1.Val + l2.Val) / 10

	for l1.Next != nil && l2.Next != nil {
		tmp := l1.Next.Val + l2.Next.Val + add
		add = tmp / 10
		head.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		l1 = l1.Next
		l2 = l2.Next
		head = head.Next
	}

	for l1.Next != nil {
		tmp := l1.Next.Val + add
		add = tmp / 10
		head.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		l1 = l1.Next
		head = head.Next
	}

	for l2.Next != nil {
		tmp := l2.Next.Val + add
		add = tmp / 10
		head.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		l2 = l2.Next
		head = head.Next
	}

	if add != 0 {
		head.Next = &ListNode{
			Val:  add,
			Next: nil,
		}
		head = head.Next
	}
	return ans
}