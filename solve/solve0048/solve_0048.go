package solve0048

/**
 * @index 48
 * @title 旋转图像
 * @difficulty 中等
 * @tags array,math,matrix
 * @draft false
 * @link https://leetcode.cn/problems/rotate-image/
 * @frontendId 48
 */

func rotate(matrix [][]int) {
	for i, n := 0, len(matrix); i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] =
				matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}
