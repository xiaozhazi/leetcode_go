/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	list := []*TreeNode{root}
	currentLevelEndMark := root
	nextLevelEndMark := root
	level := []int{}

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
		if tmp == currentLevelEndMark {
			currentLevelEndMark = nextLevelEndMark
			ans = append(ans, level)
			level = []int{}
		}
	}

	for i, j := 0, len(ans)-1; i < j; {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}