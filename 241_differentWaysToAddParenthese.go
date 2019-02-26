func calculate(x, y int, op byte) int {
	if op == '+' {
		return x + y
	} else if op == '-' {
		return x - y
	}
	return x * y
}

func diffWaysToCompute(input string) []int {
	ans := []int{}

	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			lefts := diffWaysToCompute(input[:i])
			rights := diffWaysToCompute(input[i+1:])
			for j := 0; j < len(lefts); j++ {
				for k := 0; k < len(rights); k++ {
					tmp := calculate(lefts[j], rights[k], input[i])
					ans = append(ans, tmp)
				}
			}
		}
	}
	if len(ans) == 0 {
		v, _ := strconv.Atoi(input)
		ans = append(ans, v)
	}
	return ans
}