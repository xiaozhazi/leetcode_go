/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var m map[int]bool

func searchTree(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = true
	return searchTree(root.Left, k) || searchTree(root.Right, k)
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	m = make(map[int]bool)

	return searchTree(root, k)
}