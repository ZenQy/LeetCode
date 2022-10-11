package solve0304

/**
 * @index 304
 * @title 二维区域和检索 - 矩阵不可变
 * @difficulty 中等
 * @tags design,array,matrix,prefix-sum
 * @draft false
 * @link https://leetcode.cn/problems/range-sum-query-2d-immutable/
 * @frontendId 304
 */

type NumMatrix struct {
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 && i-1 < m && j-1 >= 0 && j-1 < n {
				matrix[i][j] -= matrix[i-1][j-1]
			}
			if i-1 >= 0 && i-1 < m && j >= 0 && j < n {
				matrix[i][j] += matrix[i-1][j]
			}
			if i >= 0 && i < m && j-1 >= 0 && j-1 < n {
				matrix[i][j] += matrix[i][j-1]
			}
		}
	}
	return NumMatrix{matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := this.Matrix[row2][col2]
	if row1-1 >= 0 {
		sum -= this.Matrix[row1-1][col2]
	}
	if col1-1 >= 0 {
		sum -= this.Matrix[row2][col1-1]
	}
	if row1-1 >= 0 && col1-1 >= 0 {
		sum += this.Matrix[row1-1][col1-1]
	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
