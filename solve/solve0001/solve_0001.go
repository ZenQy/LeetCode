package solve0001

/**
 * @index 1
 * @title 两数之和
 * @difficulty 简单
 * @tags array,hash-table
 * @draft false
 * @link https://leetcode.cn/problems/two-sum/
 * @frontendId 1
 */

func twoSum(nums []int, target int) []int {
	lenth := len(nums)
	m := map[int]int{}

	for i := 0; i < lenth; i++ {
		index, ok := m[target-nums[i]]
		if ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
