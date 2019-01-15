package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}
type IntervalList []Interval

func (il IntervalList) Len() int {
	return len(il)
}

func (il IntervalList) Less(i, j int) bool {
	if il[i].Start == il[j].Start {
		return il[i].End < il[j].End
	}
	return il[i].Start < il[j].Start
}

func (il IntervalList) Swap(i, j int) {
	tmp := il[i]
	il[i] = il[j]
	il[j] = tmp
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	var intervalList IntervalList
	intervalList = intervals

	ans := []Interval{}
	sort.Sort(intervalList)
	fmt.Println(intervalList)
	currentInterval := intervalList[0]
	for i := 1; i < len(intervalList); i++ {
		if currentInterval.End < intervalList[i].Start {
			ans = append(ans, currentInterval)
			currentInterval = intervalList[i]
		} else {
			if currentInterval.End < intervalList[i].End {
				currentInterval.End = intervalList[i].End
			}
		}
	}
	ans = append(ans, currentInterval)
	return ans
}

func main() {
	var tmp Interval
	arg := []Interval{}
	tmp.Start = 1
	tmp.End = 4
	arg = append(arg, tmp)
	tmp.Start = 2
	tmp.End = 3
	arg = append(arg, tmp)
	tmp.Start = 8
	tmp.End = 10
	arg = append(arg, tmp)
	fmt.Println(merge(arg))
}
