package main

import (
	"fmt"
)

func dfs(ans *[][]int, currentSet []int, nums []int, j int) {
	*ans = append(*ans, currentSet)
	for i := j; i < len(nums); i++ {
		// add nums[i]
		currentSet = append(currentSet, nums[i])
		//recursive
		dfs(ans, currentSet, nums, i+1)
		// remove nums[i]
		size := len(currentSet)
		currentSet = currentSet[0:size]
	}
}

func subsets(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 1 {
		return ans
	}
	ptr := &ans
	tmp := []int{}
	dfs(ptr, tmp, nums, 0)
	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
