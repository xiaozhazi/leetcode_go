func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	step, i := 0, 0
	for i < len(nums) {
		if i+nums[i] >= len(nums)-1 {
			step++
			break
		}
		max := 0
		index := i + 1
		for j := i + 1; j <= nums[i]+i; j++ {
			if max < nums[j]+j {
				max = nums[j] + j
				index = j
			}
		}
		i = index
		step++
	}
	return step
}