func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	bucket := make([][]int, len(nums)+1)
	for key, value := range m {
		bucket[value] = append(bucket[value], key)
	}
	ans := []int{}

	for i, topNum := len(bucket)-1, 0; i >= 0 && topNum < k; i-- {
		if len(bucket[i]) > 0 {
			ans = append(ans, bucket[i]...)
			topNum += len(bucket[i])
		}
	}
	return ans
}