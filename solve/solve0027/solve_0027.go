package solve0027

/**
 * @index 27
 * @title 移除元素
 * @difficulty 简单
 * @tags array,two-pointers
 * @draft false
 * @link https://leetcode.cn/problems/remove-element/
 * @frontendId 27
 */

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
