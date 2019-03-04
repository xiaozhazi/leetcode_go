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
	len1, len2 := 0, 0
	for tmp := l1; tmp != nil; tmp = tmp.Next {
		len1++
	}
	for tmp := l2; tmp != nil; tmp = tmp.Next {
		len2++
	}
	if len1 < len2 {
		len1, len2, l1, l2 = len2, len1, l2, l1
	}

	array1 := make([]*ListNode, 0, len1)
	array2 := make([]*ListNode, 0, len2)
	for tmp := l1; tmp != nil; tmp = tmp.Next {
		array1 = append(array1, tmp)
	}
	for tmp := l2; tmp != nil; tmp = tmp.Next {
		array2 = append(array2, tmp)
	}

	var head ListNode
	add, tmp := 0, 0
	for i := 1; i <= len1; i++ {
		if i <= len2 {
			tmp = array1[len1-i].Val + array2[len2-i].Val + add
		} else {
			tmp = array1[len1-i].Val + add
		}
		add = tmp / 10
		node := &ListNode{
			Val:  tmp % 10,
			Next: head.Next,
		}
		head.Next = node
	}

	if add > 0 {
		node := &ListNode{
			Val:  add,
			Next: head.Next,
		}
		head.Next = node
	}
	return head.Next

}