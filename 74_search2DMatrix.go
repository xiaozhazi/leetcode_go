func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	ans := false
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		if matrix[i][0] <= target && matrix[i][col-1] >= target {
			for j := 0; j < col; j++ {
				if matrix[i][j] == target {
					ans = true
				}
			}
		}
	}
	return ans
}