package solve0344

/**
 * @index 344
 * @title 反转字符串
 * @difficulty 简单
 * @tags two-pointers,string
 * @draft false
 * @link https://leetcode.cn/problems/reverse-string/
 * @frontendId 344
 */

func reverseString(s []byte) {
	count := len(s)
	for i := 0; i < count/2; i++ {
		s[i], s[count-1-i] = s[count-1-i], s[i]
	}
}
