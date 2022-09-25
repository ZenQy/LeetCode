package solve0396

/**
 * @index 396
 * @title 旋转函数
 * @difficulty 中等
 * @tags array,math,dynamic-programming
 * @draft false
 * @link https://leetcode.cn/problems/rotate-function/
 * @frontendId 396
 */

func maxRotateFunction(nums []int) int {
	count := len(nums)
	sum := 0
	for i := 0; i < count; i++ {
		sum += nums[i]
	}
	f := 0
	for i := 0; i < count; i++ {
		f += i * nums[i]
	}

	ans := f
	for i := 1; i < count; i++ {
		f += sum - count*nums[count-i]
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
