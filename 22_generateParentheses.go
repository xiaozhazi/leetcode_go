package main

import (
	"fmt"
)

func generate(ans *[]string, left int, right int, s string) {
	if left == 0 && right == 0 {
		*ans = append(*ans, s)
	}
	if left > 0 {
		generate(ans, left-1, right, s+"(")
	}
	if right > 0 && left < right {
		generate(ans, left, right-1, s+")")
	}
}

func generateParethesis(n int) []string {
	ans := []string{}
	ptr := &ans
	generate(ptr, n, n, "")
	return ans
}

func main() {
	a := generateParethesis(3)
	fmt.Println(a)
}
