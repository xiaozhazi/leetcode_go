func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}

func nthUglyNumber(n int) int {
	nums := []int{1}
	idx := make([]int, 6)
	for i := 0; i < n; i++ {
		a, b, c := nums[idx[2]]*2, nums[idx[3]]*3, nums[idx[5]]*5
		tmp := min3(a, b, c)
		nums = append(nums, tmp)

		if tmp == a {
			idx[2]++
		}
		if tmp == b {
			idx[3]++
		}
		if tmp == c {
			idx[5]++
		}
	}
	return nums[n-1]
}