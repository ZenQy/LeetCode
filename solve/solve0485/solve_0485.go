package solve0485

/**
 * @index 485
 * @title 最大连续 1 的个数
 * @difficulty 简单
 * @tags array
 * @draft false
 * @link https://leetcode.cn/problems/max-consecutive-ones/
 * @frontendId 485
 */

func findMaxConsecutiveOnes(nums []int) int {
	max, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}
