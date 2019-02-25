func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func largestOverlap(A [][]int, B [][]int) int {
	N := len(A)
	ans := 0
	newA := []int{}
	newB := []int{}

	for i := 0; i < N*N; i++ {
		if A[i/N][i%N] == 1 {
			newA = append(newA, (i/N)*100+i%N)
		}
		if B[i/N][i%N] == 1 {
			newB = append(newB, (i/N)*100+i%N)
		}
	}
	m := make(map[int]int)

	for i := 0; i < len(newA); i++ {
		for j := 0; j < len(newB); j++ {
			m[newA[i]-newB[j]]++
			ans = max(ans, m[newA[i]-newB[j]])
		}
	}

	return ans
}