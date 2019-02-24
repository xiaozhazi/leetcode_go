func findDisappearedNumbers(nums []int) []int {
	ans := []int{}
	nums = append([]int{0}, nums...)
	for i := 1; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				break
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != i {
			ans = append(ans, i)
		}
	}
	return ans
}