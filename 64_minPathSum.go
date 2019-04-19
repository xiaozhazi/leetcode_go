func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Solution I, easy to implement
func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// Solution 2:  reduce the space

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)

	for j, _ := range dp {
		dp[j] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}