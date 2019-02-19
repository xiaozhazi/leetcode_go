
//考虑负数的左移右移
func singleNumber(nums []int) int {
	ans := 0
	negativeNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			negativeNum++
			nums[i] = 0 - nums[i]
		}
	}
	for i := uint(0); i < 32; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> i) & 1
		}
		ans |= (sum % 3) << i
	}
	if negativeNum%3 == 1 {
		return 0 - ans
	}
	return ans
}

//ab
//00
//01
//10
//00...

func singleNumber(nums []int) int {
	a := 0
	b := 0
	for i := 0; i < len(nums); i++ {
		b = (b ^ nums[i]) & ^a
		a = (a ^ nums[i]) & ^b
	}
	return b
}