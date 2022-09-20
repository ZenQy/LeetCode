package solve0118

/**
 * @index 118
 * @title 杨辉三角
 * @difficulty 简单
 * @tags array,dynamic-programming
 * @draft false
 * @link https://leetcode.cn/problems/pascals-triangle/
 * @frontendId 118
 */

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}
