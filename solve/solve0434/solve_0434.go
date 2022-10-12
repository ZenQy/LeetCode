package solve0434

/**
 * @index 434
 * @title 字符串中的单词数
 * @difficulty 简单
 * @tags string
 * @draft false
 * @link https://leetcode.cn/problems/number-of-segments-in-a-string/
 * @frontendId 434
 */

func countSegments(s string) int {
	ans, count := 0, len(s)
	for i := 0; i < count; i++ {
		if s[i] != ' ' && (i == 0 || s[i-1] == ' ') {
			ans++
		}
	}

	return ans
}
