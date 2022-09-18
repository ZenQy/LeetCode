package solve0665

/**
 * @index 665
 * @title 非递减数列
 * @difficulty 中等
 * @tags array
 * @draft false
 * @link https://leetcode.cn/problems/non-decreasing-array/
 * @frontendId 665
 */

func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
