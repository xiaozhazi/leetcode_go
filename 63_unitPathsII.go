// First use Backtracking algorithm, get TLE :(

func DFS(obstacleGrid [][]int, i, j int) {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if i == m-1 && j == n-1 && obstacleGrid[i][j] != 1 {
		count++
		return
	}
	if i < m-1 && obstacleGrid[i+1][j] == 0 {
		DFS(obstacleGrid, i+1, j)
	}
	if j < n-1 && obstacleGrid[i][j+1] == 0 {
		DFS(obstacleGrid, i, j+1)
	}
}

var count = 0

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	count = 0
	DFS(obstacleGrid, 0, 0)
	return count
}

// Then use DP to solve this issue

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 {
		return 0
	}

	m, n, count := len(obstacleGrid), len(obstacleGrid[0]), 0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return count
	}

	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				if j != 0 {
					dp[j] += dp[j-1]
				}
			}

		}
	}
	return dp[n-1]
}