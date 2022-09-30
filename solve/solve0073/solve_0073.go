package solve0073

/**
 * @index 73
 * @title 矩阵置零
 * @difficulty 中等
 * @tags array,hash-table,matrix
 * @draft false
 * @link https://leetcode.cn/problems/set-matrix-zeroes/
 * @frontendId 73
 */

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	top := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					top = true
				} else {
					matrix[i][0] = 0
					matrix[0][j] = 0
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if top {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
}
