/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	if root == nil {
		return ans
	}

	path := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{path}
	}

	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)

	path = path + "->"
	for i := 0; i < len(left); i++ {
		ans = append(ans, path+left[i])
	}
	for i := 0; i < len(right); i++ {
		ans = append(ans, path+right[i])
	}
	return ans

}