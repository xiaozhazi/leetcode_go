func hammingDistance(x int, y int) int {
	ans := 0
	cal := x ^ y
	for cal > 0 {
		cal &= cal - 1
		ans++
	}
	return ans
}