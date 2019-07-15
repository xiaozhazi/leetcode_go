package main

func twoSum(nums []int, target int) []int {
	// for i := 0; i < len(nums) - 1; i++ {
	//     for j := i + 1; j < len(nums); j++ {
	//         if nums[i] + nums[j] == target {
	//             return []int{i,j}
	//         }
	//     }
	// }
	// return []int{}

	index := map[int]int{}

	for i := 0; i < len(nums); i++ {
		index[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		value, status := index[tmp]
		if status && value != i {
			return []int{i, index[tmp]}
		}
	}
	return []int{}
}