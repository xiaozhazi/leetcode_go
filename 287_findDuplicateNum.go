func findDuplicate(nums []int) int {
	left := 1
	right := len(nums) - 1

	for left < right {
		count := 0
		mid := left + (right-left)/2
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}