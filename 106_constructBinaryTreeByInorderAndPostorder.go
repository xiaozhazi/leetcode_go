/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rightIndex := len(postorder) - 1
	root := &TreeNode{
		Val:   postorder[rightIndex],
		Left:  nil,
		Right: nil,
	}
	index := 0
	for index = 0; inorder[index] != postorder[rightIndex]; index++ {
	}
	root.Left = buildTree(inorder[:index], postorder[0:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:rightIndex])
	return root
}