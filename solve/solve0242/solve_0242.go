package solve0242

/**
 * @index 242
 * @title 有效的字母异位词
 * @difficulty 简单
 * @tags hash-table,string,sorting
 * @draft false
 * @link https://leetcode.cn/problems/valid-anagram/
 * @frontendId 242
 */

func isAnagram(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}
	cnt := [26]int{}
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if cnt[i] != 0 {
			return false
		}
	}
	return true
}
