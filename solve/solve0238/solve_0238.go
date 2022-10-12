package solve0238

/**
 * @index 238
 * @title 除自身以外数组的乘积
 * @difficulty 中等
 * @tags array,prefix-sum
 * @draft false
 * @link https://leetcode.cn/problems/product-of-array-except-self/
 * @frontendId 238
 */

func productExceptSelf(nums []int) []int {
	product, zero, count := 1, 0, len(nums)

	for i := 0; i < count; i++ {
		if nums[i] != 0 {
			product *= nums[i]
		} else {
			zero++
		}
	}

	for i := 0; i < count; i++ {
		if zero == 0 {
			nums[i] = product / nums[i]
		} else if zero > 1 {
			nums[i] = 0
		} else {
			if nums[i] == 0 {
				nums[i] = product
			} else {
				nums[i] = 0
			}
		}
	}
	return nums
}
