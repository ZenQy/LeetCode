package solve0054

/**
 * @index 54
 * @title 螺旋矩阵
 * @difficulty 中等
 * @tags array,matrix,simulation
 * @draft false
 * @link https://leetcode.cn/problems/spiral-matrix/
 * @frontendId 54
 */

func spiralOrder(matrix [][]int) []int {
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	i := 0
	count := (bottom + 1) * (right + 1)
	s := make([]int, count)

	for i < count {
		for j := left; j <= right && i < count; j++ {
			s[i] = matrix[top][j]
			i++
		}
		top++
		for j := top; j <= bottom && i < count; j++ {
			s[i] = matrix[j][right]
			i++
		}
		right--
		for j := right; j >= left && i < count; j-- {
			s[i] = matrix[bottom][j]
			i++
		}
		bottom--
		for j := bottom; j >= top && i < count; j-- {
			s[i] = matrix[j][left]
			i++
		}
		left++
	}
	return s
}
