package solve0645

/**
 * @index 645
 * @title 错误的集合
 * @difficulty 简单
 * @tags bit-manipulation,array,hash-table,sorting
 * @draft false
 * @link https://leetcode.cn/problems/set-mismatch/
 * @frontendId 645
 */

func findErrorNums(nums []int) []int {
	m := map[int]int{}
	ans := make([]int, 2)
	for _, num := range nums {
		m[num]++
	}
	for i := 1; i <= len(nums); i++ {
		if m[i] == 0 {
			ans[1] = i
		} else if m[i] == 2 {
			ans[0] = i
		}
	}
	return ans
}
