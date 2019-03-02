func getMaxMin(x, y, z int) (int, int) {
	if x < y {
		x, y = y, x
	}
	if x < z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}
	return x, y
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max, min := make([]int, len(nums)), make([]int, len(nums))

	max[0], min[0] = nums[0], nums[0]
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		tmpMax, tmpMin := max[i-1]*nums[i], min[i-1]*nums[i]
		max[i], min[i] = getMaxMin(tmpMax, tmpMin, nums[i])
		if max[i] > ans {
			ans = max[i]
		}
	}
	return ans
}