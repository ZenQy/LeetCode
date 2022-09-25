package solve0189

/**
 * @index 189
 * @title 轮转数组
 * @difficulty 中等
 * @tags array,math,two-pointers
 * @draft false
 * @link https://leetcode.cn/problems/rotate-array/
 * @frontendId 189
 */

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
