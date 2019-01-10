package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	if len(prices) < 1 {
		return max
	}

	dp := make([]int, len(prices))
	dp[0] = 0

	for i := 1; i < len(prices); i++ {
		dp[i] = dp[i-1] + prices[i] - prices[i-1]
		if dp[i] < 0 {
			dp[i] = 0
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	ans := maxProfit(prices)
	fmt.Println(ans)
}
