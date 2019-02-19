/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDepth(root *TreeNode) int {
	stack := []*TreeNode{root}
	depth := 1
	tmp := root
	last := root
	flag := root

	if root == nil {
		return 0
	}

	for len(stack) > 0 {
		tmp = stack[0]
		if tmp.Left == nil && tmp.Right == nil {
			return depth
		}
		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
			last = tmp.Left
		}
		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
			last = tmp.Right
		}
		if tmp == flag {
			flag = last
			depth++
		}
		stack = stack[1:]
	}
	return depth
}