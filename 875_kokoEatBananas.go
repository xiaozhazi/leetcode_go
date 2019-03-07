func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isFinish(piles []int, H, mid int) bool {
	if mid == 0 {
		return false
	}

	ans := 0

	for i := 0; i < len(piles); i++ {
		ans += piles[i] / mid
		if piles[i]%mid != 0 {
			ans++
		}
	}
	return ans <= H
}

func minEatingSpeed(piles []int, H int) int {
	if len(piles) > H || len(piles) < 1 {
		return -1
	}

	left, right := 0, 0

	for i := 0; i < len(piles); i++ {
		right = max(right, piles[i])
	}

	for left < right {
		mid := left + (right-left)/2
		if isFinish(piles, H, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left

}
