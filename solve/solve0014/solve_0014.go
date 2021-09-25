package solve0014

/**
 * @index 14
 * @title 最长公共前缀
 * @difficulty 简单
 * @tags string
 * @draft false
 * @link https://leetcode-cn.com/problems/longest-common-prefix/
 * @frontendId 14
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}

	minLength := len(strs[0])

	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}

	low, high := 0, minLength

	for low < high {
		mid := (high + low + 1) / 2
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return strs[0][:low]
}
