func isOccurs(array []int, n int) bool {
	if array == nil {
		return false
	}
	length := len(array)
	for i := 0; i < length; i++ {
		if array[i] == n {
			return true
		}
	}
	return false
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	occurs := make([]int, 1)
	occurs[0] = n
	for n != 1 {
		tmp := 0
		for n != 0 {
			mod := n % 10
			tmp += mod * mod
			n = n / 10
		}
		n = tmp
		if tmp == 1 {
			return true
		}
		if isOccurs(occurs, tmp) {
			return false
		}
		occurs = append(occurs, n)
	}
	return false
}