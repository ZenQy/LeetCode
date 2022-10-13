package solve0557

/**
 * @index 557
 * @title 反转字符串中的单词 III
 * @difficulty 简单
 * @tags two-pointers,string
 * @draft false
 * @link https://leetcode.cn/problems/reverse-words-in-a-string-iii/
 * @frontendId 557
 */

func reverseWords(s string) string {
	bs, count, left, right := []byte(s), len(s), 0, 0
	for i := 0; i < count; i++ {
		if bs[i] != ' ' && (i == 0 || bs[i-1] == ' ') {
			left = i
		}
		if bs[i] != ' ' && (i+1 == count || bs[i+1] == ' ') {
			right = i
			reverseBytes(bs[left : right+1])
		}

	}
	return string(bs)
}

func reverseBytes(bs []byte) {
	count := len(bs)
	for i := 0; i < count/2; i++ {
		bs[i], bs[count-1-i] = bs[count-1-i], bs[i]
	}
}
