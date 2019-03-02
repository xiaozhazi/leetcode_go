func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}