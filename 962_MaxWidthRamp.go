package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxWidthRamp(A []int) int {
	ans := 0
	stack := []int{0}

	//Create a Monotonic stack
	for i := 1; i < len(A); i++ {
		if A[stack[len(stack)-1]] > A[i] {
			stack = append(stack, i)
		} else {
			// use binary algorithm to optimization
			low := 0
			high := len(stack) - 1
			for low <= high {
				mid := (low + high) >> 1
				if A[stack[mid]] <= A[i] {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
			if low < len(stack) {
				ans = Max(ans, i-stack[low])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
}
