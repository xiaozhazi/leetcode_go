package main

func threeSum(nums []int) [][]int {
	ans := [][]int{}

	if len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			find := target - nums[j]
			if find < nums[j] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if find == nums[k] {
					tmp := []int{nums[i], nums[j], nums[k]}
					ans = append(ans, tmp)
					break
				}
			}
		}
	}
	return ans
}