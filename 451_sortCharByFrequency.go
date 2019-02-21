func frequencySort(s string) string {
	var charmaps [256]int
	for _, v := range s {
		charmaps[v]++
	}
	strmaps := make(map[int]string)

	maxIndex := 0
	for i := 0; i < 256; i++ {
		if charmaps[i] > 0 {
			strmaps[charmaps[i]] += string(i)
			if charmaps[i] > maxIndex {
				maxIndex = charmaps[i]
			}
		}
	}
	ans := ""
	for i := maxIndex; i >= 0; i-- {
		if strmaps[i] != "" {
			for _, v := range strmaps[i] {
				for j := 0; j < i; j++ {
					ans += string(v)
				}
			}
		}
	}
	return ans
}