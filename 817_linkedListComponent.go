/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	if head == nil {
		return 0
	}
	m := map[int]int{}
	ans := 0

	for i := 0; i < len(G); i++ {
		m[G[i]] = 1
	}

	for head != nil {
		if m[head.Val] == 1 {
			ans++
			for head != nil && m[head.Val] == 1 {
				head = head.Next
			}
		}
		if head != nil {
			head = head.Next
		}
	}
	return ans
}