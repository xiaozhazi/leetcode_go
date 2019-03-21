func DFS(nums, used []int, l, sum, k, target int) bool {
	if k == 0 {
		return true
	}

	if sum == target {
		return DFS(nums, used, 0, 0, k-1, target)
	}

	if sum > target {
		return false
	}

	for i := l; i < len(nums); i++ {
		if used[i] == 1 {
			continue
		}
		used[i] = 1
		if DFS(nums, used, i+1, sum+nums[i], k, target) {
			return true
		}
		used[i] = 0
	}
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	if k == 1 {
		return true
	}

	if len(nums) < k {
		return false
	}

	sum := 0
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%k != 0 {
		return false
	}

	used := make([]int, len(nums))

	return DFS(nums, used, 0, 0, k, sum/k)
}