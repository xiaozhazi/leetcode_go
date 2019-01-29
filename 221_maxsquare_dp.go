func min3(x, y, z int) int {
	if x <= y && x <= z {
		return x
	}
	if y <= z {
		return y
	}
	return z
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n, ans := len(matrix), len(matrix[0]), 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		if dp[i][0] > ans {
			ans = dp[i][0]
		}
	}

	for i := 0; i < n; i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		if dp[0][i] > ans {
			ans = dp[0][i]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans * ans
}