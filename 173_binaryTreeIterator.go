/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack []*TreeNode

type BSTIterator struct {
	stack Stack
}

func (s *Stack) Push(t *TreeNode) {
	*s = append(*s, t)
}

func (s *Stack) Pop() *TreeNode {
	if s.Len() > 0 {
		index := s.Len() - 1
		ans := (*s)[index]
		*s = (*s)[:index]
		return ans
	}
	return nil
}

func Constructor(root *TreeNode) BSTIterator {
	tmp := Stack{}
	ans := BSTIterator{stack: tmp}
	ans.PushLeft(root)
	return ans
}

func (s *Stack) Len() int {
	return len(*s)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	currentNode := this.stack.Pop()
	this.PushLeft(currentNode.Right)
	return currentNode.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stack.Len() > 0
}

func (this *BSTIterator) PushLeft(root *TreeNode) {
	for root != nil {
		this.stack.Push(root)
		root = root.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */