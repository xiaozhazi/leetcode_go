func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func maxProfit(prices []int, fee int) int {
	unhold, hold := 0, -65535
	if len(prices) <= 1 {
		return unhold
	}
	for i := 0; i < len(prices); i++ {
		tmp := unhold
		//Keep previous day status or sell previous day stock
		unhold = max(unhold, hold+prices[i])
		//Keep hold stock, or previous unhold stock and today buy some stock
		hold = max(hold, tmp-prices[i]-fee)
	}

	return unhold
}