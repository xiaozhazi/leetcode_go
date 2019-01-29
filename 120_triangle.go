func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	length, ans := len(triangle), 65536
	dp := make([][]int, len(triangle))
	for i := 0; i < length; i++ {
		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			dp[i][j] = 65536
		}
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < length; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	for i := 0; i < length; i++ {
		if dp[length-1][i] < ans {
			ans = dp[length-1][i]
		}
	}
	return ans
}