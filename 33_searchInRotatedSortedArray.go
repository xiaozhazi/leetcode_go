func binarySearch(nums []int, left, right, target int) int {
	if left >= right {
		if nums[left] == target {
			return left
		}
		return -1
	}

	if nums[left] < nums[right] {
		if nums[left] > target || nums[right] < target {
			return -1
		}
		mid := left + (right-left)/2
		if nums[mid] > target {
			return binarySearch(nums, left, mid-1, target)
		} else if nums[mid] < target {
			return binarySearch(nums, mid+1, right, target)
		} else {
			return mid
		}
	} else {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		ans := binarySearch(nums, left, mid-1, target)
		if ans == -1 {
			return binarySearch(nums, mid+1, right, target)
		}
		return ans
	}

}

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}