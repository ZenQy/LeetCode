package solve0049

/**
 * @index 49
 * @title 字母异位词分组
 * @difficulty 中等
 * @tags array,hash-table,string,sorting
 * @draft false
 * @link https://leetcode.cn/problems/group-anagrams/
 * @frontendId 49
 */

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
