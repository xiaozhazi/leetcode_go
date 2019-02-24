func setZeroes(matrix [][]int) {
	firstColHasZero := false
	firstRowHasZero := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			firstRowHasZero = true
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		if firstColHasZero {
			matrix[i][0] = 0
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if firstRowHasZero {
			matrix[0][i] = 0
		}
	}

}