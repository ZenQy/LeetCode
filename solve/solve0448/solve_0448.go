package solve0448

/**
 * @index 448
 * @title 找到所有数组中消失的数字
 * @difficulty 简单
 * @tags array,hash-table
 * @draft false
 * @link https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/
 * @frontendId 448
 */

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		nums[abs(num)-1] = -abs(nums[abs(num)-1])
	}
	ln := 0
	for i, num := range nums {
		if num > 0 {
			nums[ln] = i + 1
			ln++
		}
	}
	return nums[:ln]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
