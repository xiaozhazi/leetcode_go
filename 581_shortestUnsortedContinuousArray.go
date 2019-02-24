func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	left, right := -1, -2
	max, min := nums[0], nums[len(nums)-1]

	for i := 0; i < len(nums); i++ {
		if max <= nums[i] {
			max = nums[i]
		} else {
			right = i
		}
		if min >= nums[len(nums)-i-1] {
			min = nums[len(nums)-i-1]
		} else {
			left = len(nums) - i - 1
		}
	}
	return right - left + 1
}