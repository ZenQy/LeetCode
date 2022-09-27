package solve0059

/**
 * @index 59
 * @title 螺旋矩阵 II
 * @difficulty 中等
 * @tags array,matrix,simulation
 * @draft false
 * @link https://leetcode.cn/problems/spiral-matrix-ii/
 * @frontendId 59
 */

func generateMatrix(n int) [][]int {
	top, bottom, left, right := 0, n-1, 0, n-1
	index := 0
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
	}
	for index < n*n {
		for i := left; i <= right && index < n*n; i++ {
			index++
			s[top][i] = index
		}
		top++
		for i := top; i <= bottom && index < n*n; i++ {
			index++
			s[i][right] = index
		}
		right--
		for i := right; i >= left && index < n*n; i-- {
			index++
			s[bottom][i] = index
		}
		bottom--
		for i := bottom; i >= top && index < n*n; i-- {
			index++
			s[i][left] = index
		}
		left++
	}

	return s
}
