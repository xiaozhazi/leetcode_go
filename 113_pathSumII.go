/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans [][]int
var path []int

func DFS(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			tmp := make([]int, 0)
			tmp = append(tmp, path...)
			ans = append(ans, tmp)
		}
	}
	DFS(root.Left, sum)
	DFS(root.Right, sum)
	path = path[:len(path)-1]
}

func pathSum(root *TreeNode, sum int) [][]int {
	ans = [][]int{}
	path = []int{}
	DFS(root, sum)
	return ans
}