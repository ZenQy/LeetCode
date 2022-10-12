package solve0520

/**
 * @index 520
 * @title 检测大写字母
 * @difficulty 简单
 * @tags string
 * @draft false
 * @link https://leetcode.cn/problems/detect-capital/
 * @frontendId 520
 */

func detectCapitalUse(word string) bool {
	count, upper := len(word), 0

	for i := 0; i < count; i++ {
		if isUpper(word[i]) {
			upper++
		}
	}

	return upper == 0 || upper == count || (upper == 1 && isUpper(word[0]))

}

func isUpper(x byte) bool {
	if x >= 'A' && x <= 'Z' {
		return true
	}
	return false
}
