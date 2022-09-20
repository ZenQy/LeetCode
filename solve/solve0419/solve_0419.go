package solve0419

/**
 * @index 419
 * @title 甲板上的战舰
 * @difficulty 中等
 * @tags depth-first-search,array,matrix
 * @draft false
 * @link https://leetcode-cn.com/problems/battleships-in-a-board/
 * @frontendId 419
 */

func countBattleships(board [][]byte) int {
	count, m, n := 0, len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' &&
				(i == 0 || board[i-1][j] == '.') &&
				(j == 0 || board[i][j-1] == '.') {
				count++
			}
		}
	}

	return count
}
