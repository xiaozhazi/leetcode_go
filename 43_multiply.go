package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	a := make([]int, len(num1))
	b := make([]int, len(num2))
	ans := make([]int, len(num1)+len(num2))

	for index, ch := range num1 {
		a[index] = int(ch - '0')
	}

	for index, ch := range num2 {
		b[index] = int(ch - '0')
	}

	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			ans[i+j+1] += a[i] * b[j]
		}
	}

	ansStr := ""
	for k := len(num1) + len(num2) - 1; k >= 0; k-- {
		if k > 0 {
			ans[k-1] += ans[k] / 10
		}
		ans[k] %= 10
		ansStr = string(ans[k]+'0') + ansStr
	}

	for len(ansStr) > 1 && ansStr[0] == '0' {
		ansStr = ansStr[1:]
	}
	return ansStr
}

func main() {
	fmt.Println(multiply("123", "456"))
}
