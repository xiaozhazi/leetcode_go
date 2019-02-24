func firstMissingPositive(nums []int) int {
	recNums := make([]int, len(nums)+2)
	for _, val := range nums {
		if val > len(nums) || val <= 0 {
			continue
		}
		recNums[val] = 1
	}
	for i := 1; i <= len(nums); i++ {
		if recNums[i] == 0 {
			return i
		}
	}
	return len(nums) + 1
}