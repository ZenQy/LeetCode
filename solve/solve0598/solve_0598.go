package solve0598

/**
 * @index 598
 * @title 范围求和 II
 * @difficulty 简单
 * @tags array,math
 * @draft false
 * @link https://leetcode-cn.com/problems/range-addition-ii/
 * @frontendId 598
 */

func maxCount(m int, n int, ops [][]int) int {
	for i := 0; i < len(ops); i++ {
		m, n = min(m, ops[i][0]), min(ops[i][1], n)
	}

	return m * n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
