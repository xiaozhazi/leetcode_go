func getNeighborLive(board [][]int, x, y int) int {
	row := len(board)
	col := len(board[0])
	count := 0

	for i := -1; i != 2; i++ {
		for j := -1; j != 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			xn := x + i
			yn := y + j
			if xn >= 0 && xn < row && yn >= 0 && yn < col && (board[xn][yn] == 1 || board[xn][yn] == 2) {
				count++
			}
		}
	}
	return count
}

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			count := getNeighborLive(board, i, j)
			if board[i][j] == 1 {
				if count < 2 || count > 3 {
					board[i][j] = 2
				}
			} else {
				if count == 3 {
					board[i][j] = 3
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}