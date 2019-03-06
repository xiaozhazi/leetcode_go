func binarySearch(nums []int, left, right, target int) bool {
	if left >= right {
		if nums[left] == target {
			return true
		}
		return false
	}

	if nums[left] < nums[right] {
		mid := left + (right-left)/2
		if nums[mid] < target {
			return binarySearch(nums, mid+1, right, target)
		} else if nums[mid] > target {
			return binarySearch(nums, left, mid-1, target)
		}
		return true
	} else {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		ans := binarySearch(nums, left, mid-1, target)
		if !ans {
			return binarySearch(nums, mid+1, right, target)
		}
		return ans
	}
}

func search(nums []int, target int) bool {
	if len(nums) < 1 {
		return false
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}