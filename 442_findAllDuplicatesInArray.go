func findDuplicates(nums []int) []int {
	ans := []int{}

	if len(nums) < 1 {
		return ans
	}
	nums = append([]int{0}, nums...)
	set := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				if _, ok := set[nums[i]]; ok == false {
					ans = append(ans, nums[i])
					set[nums[i]] = false
				}
				break
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

func findDuplicates(nums []int) []int {
	ans := []int{}

	if len(nums) < 1 {
		return ans
	}
	// Use + - to distinguish
	for i := 0; i < len(nums); i++ {
		nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		if nums[abs(nums[i])-1] > 0 {
			ans = append(ans, abs(nums[i]))
		}
	}
	return ans
}