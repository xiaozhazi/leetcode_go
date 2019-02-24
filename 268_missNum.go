func missingNumber(nums []int) int {
	misNum := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	misNum = len(nums)*(len(nums)+1)/2 - sum
	return misNum
}