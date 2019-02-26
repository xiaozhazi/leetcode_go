/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generate(start, end int) []*TreeNode {
	var ans []*TreeNode
	if start > end {
		return []*TreeNode{nil}
	}

	for i := start; i <= end; i++ {
		leftTrees := generate(start, i-1)
		rightTrees := generate(i+1, end)
		for j := 0; j < len(leftTrees); j++ {
			for k := 0; k < len(rightTrees); k++ {
				root := &TreeNode{
					Val:   i,
					Left:  leftTrees[j],
					Right: rightTrees[k],
				}
				ans = append(ans, root)
			}
		}
	}
	return ans
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}