func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func increasingTriplet(nums []int) bool {
	ans := 0
	length := len(nums)

	if length == 0 {
		return false
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
	if ans >= 3 {
		return true
	}
	return false
}