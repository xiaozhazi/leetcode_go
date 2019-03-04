func reverse(nums []int) {
	for left, right := 0, len(nums)-1; left < right; {
		nums[left], nums[right] = nums[right], nums[left]
		right--
		left++
	}
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	index := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i - 1
			break
		}
	}
	if index == -1 {
		reverse(nums)
		return
	}

	bigIndex := -1

	for bigIndex = len(nums) - 1; bigIndex > index; bigIndex-- {
		if nums[bigIndex] > nums[index] {
			break
		}
	}
	nums[bigIndex], nums[index] = nums[index], nums[bigIndex]
	s := nums[index+1 : len(nums)]
	reverse(s)
}