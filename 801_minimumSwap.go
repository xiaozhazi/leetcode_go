func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSwap(A []int, B []int) int {
	length := len(A)
	if length <= 1 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = 65536
		}
	}

	dp[0][0] = 0
	dp[0][1] = 1

	for i := 1; i < length; i++ {
		if A[i-1] < A[i] && B[i-1] < B[i] {
			dp[i][0] = min(dp[i-1][0], dp[i][0])
			dp[i][1] = min(dp[i-1][1]+1, dp[i][1])
		}

		if A[i-1] < B[i] && B[i-1] < A[i] {
			dp[i][0] = min(dp[i-1][1], dp[i][0])
			dp[i][1] = min(dp[i-1][0]+1, dp[i][1])
		}
	}
	return min(dp[length-1][0], dp[length-1][1])
}

//O(N), Space(1)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSwap(A []int, B []int) int {
	length := len(A)
	if length <= 1 {
		return 0
	}
	preNatural := 0
	preSwap := 1

	for i := 1; i < length; i++ {
		curNatural := 65536
		curSwap := 65536
		//keep order OR change both
		if A[i-1] < A[i] && B[i-1] < B[i] {
			curNatural = min(preNatural, curNatural)
			curSwap = min(preSwap+1, curSwap)
		}

		// Only can change one position
		if A[i-1] < B[i] && B[i-1] < A[i] {
			curNatural = min(preSwap, curNatural)
			curSwap = min(preNatural+1, curSwap)
		}
		// Move to next
		preNatural = curNatural
		preSwap = curSwap
	}
	return min(preNatural, preSwap)
}