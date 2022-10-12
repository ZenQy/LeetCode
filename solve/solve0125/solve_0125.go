package solve0125

import (
	"strings"
)

/**
 * @index 125
 * @title 验证回文串
 * @difficulty 简单
 * @tags two-pointers,string
 * @draft false
 * @link https://leetcode.cn/problems/valid-palindrome/
 * @frontendId 125
 */

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	count := len(s)
	for i, j := 0, count-1; i < j; {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			i++
			continue
		}
		if (s[j] < 'a' || s[j] > 'z') && (s[j] < '0' || s[j] > '9') {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
