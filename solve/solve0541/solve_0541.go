package solve0541

/**
 * @index 541
 * @title 反转字符串 II
 * @difficulty 简单
 * @tags two-pointers,string
 * @draft false
 * @link https://leetcode.cn/problems/reverse-string-ii/
 * @frontendId 541
 */

func reverseStr(s string, k int) string {
	bs := []byte(s)
	count := len(bs)
	for i := 0; i < count; i += 2 * k {
		left, right := i, min(i+k, count)-1
		for left < right {
			bs[left], bs[right] = bs[right], bs[left]
			left++
			right--
		}
	}
	return string(bs)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
