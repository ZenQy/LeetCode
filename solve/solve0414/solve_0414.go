package solve0414

import "math"

/**
 * @index 414
 * @title 第三大的数
 * @difficulty 简单
 * @tags array,sorting
 * @draft false
 * @link https://leetcode-cn.com/problems/third-maximum-number/
 * @frontendId 414
 */

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if a > num && num > b {
			b, c = num, b
		} else if b > num && num > c {
			c = num
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
