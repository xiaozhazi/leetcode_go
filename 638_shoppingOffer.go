func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isValid(needs []int, special []int) bool {
	for i := 0; i < len(needs); i++ {
		if needs[i] < special[i] {
			return false
		}
	}
	return true
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	minPrice := 0
	for i := 0; i < len(needs); i++ {
		minPrice += price[i] * needs[i]
	}

	for i := 0; i < len(special); i++ {
		if isValid(needs, special[i]) {
			currentNeeds := []int{}
			for j := 0; j < len(needs); j++ {
				currentNeeds = append(currentNeeds, needs[j]-special[i][j])
			}
			tmpPrice := shoppingOffers(price, special, currentNeeds) + special[i][len(needs)]
			minPrice = min(minPrice, tmpPrice)
		}
	}
	return minPrice
}