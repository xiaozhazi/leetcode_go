func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func max3(x, y, z int) int {
	if x > y && x > z {
		return x
	}
	if y > z {
		return y
	}
	return z
}

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	//dp[i] means rob [0...i] hourses

	//consider rob [0, length-2]
	dp := make([]int, length-1)
	dp[0] = nums[0]
	for i := 1; i < length-1; i++ {
		dp[i] = 0
		for j := 0; j <= i; j++ {
			if j >= 2 {
				dp[i] = max(dp[i], dp[j-2]+nums[j])
			} else {
				dp[i] = max(dp[i], nums[j])
			}
		}
	}

	//consider rob [1...length-1]
	dp2 := make([]int, length)
	dp2[1] = nums[1]
	for i := 2; i < length; i++ {
		dp2[i] = 0
		for j := 1; j <= i; j++ {
			if j >= 3 {
				dp2[i] = max(dp2[i], dp2[j-2]+nums[j])
			} else {
				dp2[i] = max(dp2[i], nums[j])
			}
		}
	}

	return max(dp[length-2], dp2[length-1])

}