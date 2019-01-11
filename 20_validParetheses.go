package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	ans := true
	if len(s) < 1 {
		return ans
	}
	stackList := list.New()
	stackList.PushBack(' ')
	for _, char := range s {
		back := stackList.Back()
		if (back.Value == '(' && char == ')') ||
			(back.Value == '[' && char == ']') ||
			(back.Value == '{' && char == '}') {
			stackList.Remove(back)
		} else {
			stackList.PushBack(char)
		}
	}
	if stackList.Len() != 1 {
		ans = false
	}
	return ans
}

func main() {
	s := "()[]{[}]"
	if isValid(s) {
		fmt.Println("valid!")
	} else {
		fmt.Println("invalid!")
	}
}
