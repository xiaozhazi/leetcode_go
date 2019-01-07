package main

import(
"fmt"
)

func searchRange(nums []int, target int) ([2]int){
	output := [2]int{-1, -1}
	length := len(nums)
	if length <=0 || nums[0] > target || nums[length - 1] < target {
		return output
	}
	if length == 1 && nums[0] == target{
		output[0] = 0
		output[1] = 0
		return output
	} else if length > 1{
		mid := length / 2
	    if nums[mid] < target {
		  	output = searchRange(nums[mid:length], target)
			if output[0] != -1 {
				output[0] += mid
			}
			if output[1] != -1 {
				output[1] += mid
			} 
	    } else if nums[mid] > target {
			output = searchRange(nums[0:mid], target)
			
		} else {
			output[0] = searchRange(nums[0:mid], target)[0]
			if output[0] == -1 {
			output[0] = mid
			}
			output[1] = mid + searchRange(nums[mid:length], target)[1]
		}
	}
	
	fmt.Println(output)
	return output

}



func main () {
	nums := []int{5,7,7,8,8,10}
	searchRange(nums, 6)
}