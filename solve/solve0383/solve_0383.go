package solve0383

/**
 * @index 383
 * @title 赎金信
 * @difficulty 简单
 * @tags hash-table,string,counting
 * @draft false
 * @link https://leetcode.cn/problems/ransom-note/
 * @frontendId 383
 */

func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for i := 0; i < len(magazine); i++ {
		cnt[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		cnt[ransomNote[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if cnt[i] < 0 {
			return false
		}
	}
	return true
}
