package solve0453

import "math"

/**
 * @index 453
 * @title 最小操作次数使数组元素相等
 * @difficulty 中等
 * @tags array,math
 * @draft false
 * @link https://leetcode.cn/problems/minimum-moves-to-equal-array-elements/
 * @frontendId 453
 */

func minMoves(nums []int) int {
	min := math.MaxInt
	count := len(nums)
	ans := 0
	for i := 0; i < count; i++ {
		ans += nums[i]
		if min > nums[i] {
			min = nums[i]
		}
	}
	ans -= count * min
	return ans
}
