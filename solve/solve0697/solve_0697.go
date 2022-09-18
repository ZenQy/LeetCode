package solve0697

/**
 * @index 697
 * @title 数组的度
 * @difficulty 简单
 * @tags array,hash-table
 * @draft false
 * @link https://leetcode.cn/problems/degree-of-an-array/
 * @frontendId 697
 */

func findShortestSubArray(nums []int) int {
	m := map[int][]int{}
	for i, num := range nums {
		if _, ok := m[num]; ok {
			m[num][2] = i
		} else {
			m[num] = make([]int, 3)
			m[num][1], m[num][2] = i, i
		}
		m[num][0]++
	}
	repeat, lenth := 0, 0

	for _, v := range m {
		if v[0] > repeat {
			repeat = v[0]
			lenth = v[2] - v[1]
		} else if v[0] == repeat {
			if v[2]-v[1] < lenth {
				lenth = v[2] - v[1]
			}
		}
	}

	return lenth + 1
}
