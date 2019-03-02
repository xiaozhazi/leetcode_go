func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	left := 0
	lastOccur := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if lastOccur[s[i]] == 0 || lastOccur[s[i]] < left {
			ans = max(ans, i-left+1)
		} else {
			left = lastOccur[s[i]]
		}
		lastOccur[s[i]] = i + 1
	}

	return ans
}