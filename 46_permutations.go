package main

import (
	"fmt"
	"sort"
)

var ans [][]int

func reverse(nums []int) {
	length := len(nums)
	for i := 0; i < length && i < length-i; i++ {
		tmp := nums[i]
		nums[i] = nums[length-i-1]
		nums[length-i-1] = tmp
	}
}

func hasNextPerm(nums []int) bool {
	length := len(nums)
	if len(nums) < 2 {
		return false
	}
	var i, k int
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	// When nums are desc sorted
	if i < 0 {
		return false
	}

	for k = length - 1; k > i; k-- {
		if nums[k] > nums[i] {
			//swap nums[i] and nums[k]
			tmp := nums[i]
			nums[i] = nums[k]
			nums[k] = tmp
			//then reverse nums from i+1 to end
			s := nums[i+1 : len(nums)]
			reverse(s)
			return true
		}
	}

	return false

}

func permutation(nums []int) [][]int {
	ans = [][]int{}
	sort.Ints(nums)
	// add nums into ans
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	ans = append(ans, tmp)
	// If has next permutation, add next permutation into ans
	for {
		if !hasNextPerm(nums) {
			break
		} else {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			ans = append(ans, tmp)
		}
	}
	return ans
}

func main() {
	nums := []int{0, -1, 1}
	ans = permutation(nums)
	fmt.Println(ans)
}
