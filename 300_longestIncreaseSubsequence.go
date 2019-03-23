func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLIS(nums []int) int {
	ans := 0
	length := len(nums)

	if length == 0 {
		return ans
	}

	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	ans = 1

	for i := len(nums) - 2; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			ans = max(ans, dp[i])
		}
	}
	return ans
}