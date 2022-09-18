package solve0283

/**
 * @index 283
 * @title 移动零
 * @difficulty 简单
 * @tags array,two-pointers
 * @draft false
 * @link https://leetcode.cn/problems/move-zeroes/
 * @frontendId 283
 */

func moveZeroes(nums []int) []int {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
		right++
	}
	return nums
}
