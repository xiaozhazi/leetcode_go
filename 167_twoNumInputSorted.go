func twoSum(numbers []int, target int) []int {
	ans := []int{}
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > target {
			return ans
		}
		find := target - numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] > find {
				break
			}
			if numbers[j] == find {
				return []int{i + 1, j + 1}
			}
		}
	}
	return ans
}

// 二分查找优化

func twoSum(numbers []int, target int) []int {
	ans := []int{}
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > target {
			return ans
		}
		find := target - numbers[i]
		for left, right := i+1, len(numbers)-1; left <= right; {
			mid := right + (left-right)/2
			if numbers[mid] == find {
				return []int{i + 1, mid + 1}
			}
			if numbers[mid] > find {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return ans
}