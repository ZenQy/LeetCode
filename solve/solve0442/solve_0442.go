package solve0442

/**
 * @index 442
 * @title 数组中重复的数据
 * @difficulty 中等
 * @tags array,hash-table
 * @draft false
 * @link https://leetcode.cn/problems/find-all-duplicates-in-an-array/
 * @frontendId 442
 */

func findDuplicates(nums []int) []int {
	ln := len(nums)
	for i := 0; i < ln; i++ {
		nums[i] -= 1
		nums[nums[i]%ln] += ln
	}
	right := 0
	for i := 0; i < ln; i++ {
		if nums[i]/ln == 2 {
			nums[right] = i + 1
			right++
		}
	}

	return nums[:right]
}
