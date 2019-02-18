func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = 65536
	}

	for i := 0; i*i <= n; i++ {
		dp[i*i] = 1
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j*j < i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}