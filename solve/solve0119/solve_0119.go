package solve0119

/**
 * @index 119
 * @title 杨辉三角 II
 * @difficulty 简单
 * @tags array,dynamic-programming
 * @draft false
 * @link https://leetcode.cn/problems/pascals-triangle-ii/
 * @frontendId 119
 */

func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			ans[j] += ans[j-1]
		}
	}

	return ans
}
