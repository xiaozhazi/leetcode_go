/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxCount int

func calMap(root *TreeNode, countMap map[int]int) {
	if root == nil {
		return
	}
	countMap[root.Val]++
	if maxCount < countMap[root.Val] {
		maxCount = countMap[root.Val]
	}
	calMap(root.Left, countMap)
	calMap(root.Right, countMap)
}
func findMode(root *TreeNode) []int {
	maxCount = 0
	countMap := make(map[int]int)

	calMap(root, countMap)
	ans := []int{}
	for k, v := range countMap {
		if v == maxCount {
			ans = append(ans, k)
		}
	}
	return ans
}