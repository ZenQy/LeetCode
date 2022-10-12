package solve0058

/**
 * @index 58
 * @title 最后一个单词的长度
 * @difficulty 简单
 * @tags string
 * @draft false
 * @link https://leetcode.cn/problems/length-of-last-word/
 * @frontendId 58
 */

func lengthOfLastWord(s string) int {
	ans, count := 0, len(s)
	for i := count - 1; i >= 0; i-- {
		if s[i] == ' ' && ans == 0 {
			continue
		}

		if s[i] != ' ' {
			ans++
		} else {
			return ans
		}

	}
	return ans
}
