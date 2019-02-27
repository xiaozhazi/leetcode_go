/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	ans := []float64{}
	if root == nil {
		return ans
	}

	list := []*TreeNode{root}
	curLevelEnd := root
	nextLevelEnd := root
	levelCount := 0
	levelSum := 0
	for len(list) != 0 {
		if list[0] == nil {
			list = list[1:]
			continue
		}
		tmp := list[0]
		list = list[1:]
		levelCount++
		levelSum += tmp.Val

		if tmp.Left != nil {
			nextLevelEnd = tmp.Left
			list = append(list, tmp.Left)
		}
		if tmp.Right != nil {
			nextLevelEnd = tmp.Right
			list = append(list, tmp.Right)
		}
		if tmp == curLevelEnd {
			curLevelEnd = nextLevelEnd
			average := float64(levelSum) / float64(levelCount)
			ans = append(ans, average)
			levelCount = 0
			levelSum = 0
		}
	}
	return ans
}