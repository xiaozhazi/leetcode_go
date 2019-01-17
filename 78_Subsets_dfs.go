package main

import (
	"fmt"
)

func dfs(ans *[][]int, tmp []int, nums []int, j int) {
	tmp1 := make([]int, len(tmp))
	copy(tmp1, tmp)
	*ans = append(*ans, tmp1)
	for i := j; i < len(nums); i++ {
		// add nums[i]
		tmp = append(tmp, nums[i])
		//recursive
		dfs(ans, tmp, nums, i+1)
		// remove nums[i]
		size := len(tmp)
		tmp = tmp[0 : size-1]
	}
}

func subsets(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 1 {
		return ans
	}
	ptr := &ans
	dfs(ptr, []int{}, nums, 0)
	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
