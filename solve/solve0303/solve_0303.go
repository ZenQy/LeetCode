package solve0303

/**
 * @index 303
 * @title 区域和检索 - 数组不可变
 * @difficulty 简单
 * @tags design,array,prefix-sum
 * @draft false
 * @link https://leetcode.cn/problems/range-sum-query-immutable/
 * @frontendId 303
 */

type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{nums}
}

func (na *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return na.Nums[right]
	}
	return na.Nums[right] - na.Nums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
