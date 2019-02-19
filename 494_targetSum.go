var ans int

func DFS(nums []int, index int, remain int) {
	if index == len(nums) {
		if remain == 0 {
			ans++
		}
		return
	}
	DFS(nums, index+1, remain-nums[index])
	DFS(nums, index+1, remain+nums[index])
}
func findTargetSumWays(nums []int, S int) int {
	ans = 0
	DFS(nums, 0, S)
	return ans
}