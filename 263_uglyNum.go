func isUgly(num int) bool {
	if num <= 1 {
		return num == 1
	}
	flag := true

	for num > 1 && flag {
		flag = false
		if num%2 == 0 {
			num /= 2
			flag = true
		}
		if num%3 == 0 {
			num /= 3
			flag = true
		}
		if num%5 == 0 {
			num /= 5
			flag = true
		}
	}
	return flag
}