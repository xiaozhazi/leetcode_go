func productExceptSelf(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}

	r := make([]int, len(nums))
	sum := 1
	r[0] = 1

	for i := 1; i < len(nums); i++ {
		sum *= nums[i-1]
		r[i] = sum
	}

	sum = nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		r[i] *= sum
		sum *= nums[i]
	}
	return r
}