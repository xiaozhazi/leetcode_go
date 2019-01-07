package main

import (
	"fmt"
	"sort"
)

var ans = [][]int{}
var current = []int{}

func getCombination(cand []int, require int, index int) {
	if 0 == require {
		tmp := make([]int, len(current))
		copy(tmp, current)
		ans = append(ans, tmp)
		return
	} else if require < 0 {
		return
	}

	for i := index; i < len(cand); i++ {
		current = append(current, cand[i])
		getCombination(cand, require-cand[i], i)
		current = current[:len(current)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	ans = [][]int{}
	current = []int{}
	sort.Ints(candidates)
	getCombination(candidates, target, 0)
	return ans
}

func main() {
	var candidates = []int{42, 26, 36, 38, 35, 41, 20, 47, 45, 23, 33, 39, 25, 43, 29, 31, 28, 48, 21, 46, 22, 30, 37, 32, 44, 40}
	var target = 55
	fmt.Println(target)
	fmt.Println(candidates)

	var result = [][]int{}

	result = combinationSum(candidates, target)
	fmt.Println(result)

}
