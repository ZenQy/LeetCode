package solve0389

/**
 * @index 389
 * @title 找不同
 * @difficulty 简单
 * @tags bit-manipulation,hash-table,string,sorting
 * @draft false
 * @link https://leetcode.cn/problems/find-the-difference/
 * @frontendId 389
 */

func findTheDifference(s, t string) byte {
	sum, count := byte(0), len(s)
	for i := 0; i < count; i++ {
		sum += t[i] - s[i]
	}
	sum += t[count]
	return sum
}
