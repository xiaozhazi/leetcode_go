/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	if p.Val > q.Val {
		tmp := p
		p = q
		q = tmp
	}

	if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}