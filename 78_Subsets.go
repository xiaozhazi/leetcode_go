package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) < 1 {
		return ans
	}

	max := 1 << uint16(len(nums))

	for i := 0; i < max; i++ {
		bit := i
		tmpSet := []int{}
		for j := 0; j < len(nums); j++ {
			if bit&1 == 1 {
				tmpSet = append(tmpSet, nums[j])
			}
			bit >>= 1
		}
		ans = append(ans, tmpSet)
	}
	return ans
}

func main() {
	/**
	*[]     [1]  [2] [1,2]  [3] [1,3] [2,3] [1,2,3]
	*000   001   010  011   100  101   110  111
	 */
	fmt.Println(subsets([]int{1, 2, 3}))
}
