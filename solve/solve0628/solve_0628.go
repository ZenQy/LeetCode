package solve0628

import "math"

/**
 * @index 628
 * @title 三个数的最大乘积
 * @difficulty 简单
 * @tags array,math,sorting
 * @draft false
 * @link https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
 * @frontendId 628
 */

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range nums {
		if num < min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}

		if num > max1 {
			max1, max2, max3 = num, max1, max2
		} else if num > max2 {
			max2, max3 = num, max2
		} else if num > max3 {
			max3 = num
		}
	}

	return max(min1*min2*max1, max1*max2*max3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
