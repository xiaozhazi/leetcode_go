func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	left, index, right := 0, 0, len(nums)-1

	for index <= right {
		if nums[index] == 0 {
			nums[index], nums[left] = nums[left], nums[index]
			index++
			left++
		} else if nums[index] == 2 {
			nums[index], nums[right] = nums[right], nums[index]
			right--
		} else {
			index++
		}
	}
}