func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	for i := 0; i <= amount; i++ {
		if dp[i] < 0 {
			continue
		}
		for j := 0; j < len(coins); j++ {
			if i+coins[j] > amount {
				continue
			}
			if dp[i+coins[j]] < 0 || dp[i+coins[j]] > dp[i]+1 {
				dp[i+coins[j]] = dp[i] + 1
			}
		}
	}
	return dp[amount]
}