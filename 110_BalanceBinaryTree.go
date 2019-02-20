/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func DFS(root *TreeNode, ans *bool) int {
	if *ans == false || root == nil {
		return 0
	}
	left := DFS(root.Left, ans)
	right := DFS(root.Right, ans)
	if left-right > 1 || right-left > 1 {
		*ans = false
		return 0
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func isBalanced(root *TreeNode) bool {
	ans := true
	DFS(root, &ans)
	return ans
}