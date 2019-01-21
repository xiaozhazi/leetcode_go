package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	ans := 0
	flag := true

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend < 0 && divisor < 0 {
		dividend = 0 - dividend
		divisor = 0 - divisor
	} else if dividend < 0 && divisor > 0 {
		dividend = 0 - dividend
		flag = false
	} else if dividend > 0 && divisor < 0 {
		divisor = 0 - divisor
		flag = false
	}

	for dividend >= divisor {
		tmp := divisor
		multiple := 1
		for dividend >= (tmp << 1) {
			tmp <<= 1
			multiple <<= 1
		}
		dividend -= tmp
		ans += multiple
	}

	if flag {
		return ans
	} else {
		return 0 - ans
	}

}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(-1, -1))
}
