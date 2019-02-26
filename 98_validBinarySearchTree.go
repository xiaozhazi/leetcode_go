/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var last *TreeNode

func isValid(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isValid(root.Left) == false {
		return false
	}
	if last != nil && last.Val >= root.Val {
		return false
	}
	last = root
	return isValid(root.Right)

}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	last = nil
	return isValid(root)
}