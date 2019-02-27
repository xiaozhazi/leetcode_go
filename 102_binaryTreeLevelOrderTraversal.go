/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	list := []*TreeNode{root}
	level := []int{}
	currentLevelEndMark := root
	nextLevelEndMark := root
	for len(list) != 0 {
		if list[0] == nil {
			list = list[1:]
			continue
		}
		tmp := list[0]
		list = list[1:]
		level = append(level, tmp.Val)
		if tmp.Left != nil {
			nextLevelEndMark = tmp.Left
			list = append(list, tmp.Left)
		}
		if tmp.Right != nil {
			nextLevelEndMark = tmp.Right
			list = append(list, tmp.Right)
		}
		if currentLevelEndMark == tmp {
			currentLevelEndMark = nextLevelEndMark
			cpy := make([]int, len(level))
			copy(cpy, level)
			ans = append(ans, cpy)
			level = []int{}
		}
	}
	return ans
}