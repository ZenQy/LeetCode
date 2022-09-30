package solve0566

/**
 * @index 566
 * @title 重塑矩阵
 * @difficulty 简单
 * @tags array,matrix,simulation
 * @draft false
 * @link https://leetcode.cn/problems/reshape-the-matrix/
 * @frontendId 566
 */

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	ans := make([][]int, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
		for j := 0; j < c; j++ {
			index := i*c + j
			ans[i][j] = mat[index/n][index%n]
		}
	}

	return ans
}
