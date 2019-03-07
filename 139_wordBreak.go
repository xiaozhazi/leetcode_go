func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = false
	}
	dp[0] = true

	for i := 0; i <= len(s); i++ {
		if dp[i] {
			for j := 1; i+j <= len(s); j++ {
				for k := 0; k < len(wordDict); k++ {
					if s[i:i+j] == wordDict[k] {
						dp[i+j] = true
						break
					}
				}
			}
		}
	}
	return dp[len(s)]
}