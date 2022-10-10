package solve0289

/**
 * @index 289
 * @title 生命游戏
 * @difficulty 中等
 * @tags array,matrix,simulation
 * @draft false
 * @link https://leetcode.cn/problems/game-of-life/
 * @frontendId 289
 */

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := 0

			for k := 0; k < 8; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if x >= 0 && x < m && y >= 0 && y < n && board[x][y] > 0 {
					live++
				}
			}

			if board[i][j] == 1 && (live < 2 || live > 3) {
				board[i][j] = 2
			}
			if board[i][j] == 0 && live == 3 {
				board[i][j] = -1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == -1 {
				board[i][j] = 1
			}
			if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
}
