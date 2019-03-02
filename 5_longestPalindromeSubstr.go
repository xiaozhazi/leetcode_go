// Use Manacher algorithm

func makeNewStr(s string) string {
	sLen := len(s)
	ans := make([]byte, 2*sLen+1)
	index := 0

	for i := 0; i < len(ans); i++ {
		if (i & 1) == 0 {
			ans[i] = '#'
		} else {
			ans[i] = s[index]
			index++
		}
	}
	return string(ans)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	newStr := makeNewStr(s)
	r := make([]int, len(newStr))
	R := -1
	C := -1
	maxLen := 0
	maxCenter := 0

	for i := 0; i < len(newStr); i++ {
		if R > i {
			r[i] = min(r[2*C-i], R-i)
		} else {
			r[i] = 1
		}
		for i+r[i] < len(newStr) && i-r[i] >= 0 {
			if newStr[i+r[i]] == newStr[i-r[i]] {
				r[i]++
			} else {
				break
			}
		}
		if i+r[i] > R {
			R = i + r[i]
			C = i
		}
		if r[i] > maxLen {
			maxLen = r[i]
			maxCenter = i
		}
	}

	ans := []byte{}
	fmt.Print(maxCenter)

	for i := maxCenter - maxLen + 1; i < maxCenter+maxLen; i++ {
		if newStr[i] != '#' {
			ans = append(ans, newStr[i])
		}
	}
	return string(ans)
}