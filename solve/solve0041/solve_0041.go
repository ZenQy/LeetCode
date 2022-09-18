package solve0041

/**
 * @index 41
 * @title 缺失的第一个正数
 * @difficulty 困难
 * @tags array,hash-table
 * @draft false
 * @link https://leetcode.cn/problems/first-missing-positive/
 * @frontendId 41
 */

func firstMissingPositive(nums []int) int {
	ln := len(nums)
	for i := 0; i < ln; i++ {
		for nums[i] >= 1 && nums[i] <= ln &&  nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < ln; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return ln + 1
}
