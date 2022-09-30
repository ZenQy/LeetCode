package solve0498

/**
 * @index 498
 * @title 对角线遍历
 * @difficulty 中等
 * @tags array,matrix,simulation
 * @draft false
 * @link https://leetcode.cn/problems/diagonal-traverse/
 * @frontendId 498
 */

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, m*n)
	index := 0
	for i := 0; i < m+n-1; i++ {
		if i%2 == 1 {
			y := min(i, n-1)
			x := i - y
			for x < m && y >= 0 {
				ans[index] = mat[x][y]
				index++
				x++
				y--
			}
		} else {
			x := min(i, m-1)
			y := i - x
			for x >= 0 && y < n {
				ans[index] = mat[x][y]
				index++
				x--
				y++
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
