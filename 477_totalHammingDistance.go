func getHammingDistance(x, y int) int {
	ans := 0
	cal := x ^ y
	for cal > 0 {
		cal &= cal - 1
		ans++
	}
	return ans
}

// func totalHammingDistance(nums []int) int {
//     ans := 0
//     if len(nums) < 1 {
//         return ans
//     }
//     for i := 0; i < len(nums); i++ {
//         for j := i + 1; j< len(nums); j++ {
//             ans += getHammingDistance(nums[i], nums[j])
//         }
//     }
//     return ans
// }

func totalHammingDistance(nums []int) int {
	ans := 0
	bitCount := [2]int{0, 0}
	zeroCount := 0

	for zeroCount < len(nums) {
		zeroCount = 0
		bitCount[0], bitCount[1] = 0, 0
		for i := 0; i < len(nums); i++ {
			if nums[i]&1 == 1 {
				bitCount[1]++
			} else {
				bitCount[0]++
			}
			if nums[i] == 0 {
				zeroCount++
			} else {
				nums[i] = nums[i] >> 1
			}
		}
		ans += bitCount[0] * bitCount[1]
	}
	return ans
}
