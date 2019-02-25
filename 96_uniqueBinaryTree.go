func numTrees(n int) int {
	ans := []int{1, 1, 2}
	for i := 3; i <= n; i++ {
		tmp := 0
		for j := 1; j <= i; j++ {
			tmp += ans[j-1] * ans[i-j]
		}
		ans = append(ans, tmp)
	}
	return ans[n]
}