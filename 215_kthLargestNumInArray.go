func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k == 0 {
		return 0
	}
	for i := 0; i < k; i++ {
		maxIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[maxIndex] < nums[j] {
				maxIndex = j
			}
		}
		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
	}
	return nums[k-1]
}