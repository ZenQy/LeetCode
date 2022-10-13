package solve0387

/**
 * @index 387
 * @title 字符串中的第一个唯一字符
 * @difficulty 简单
 * @tags queue,hash-table,string,counting
 * @draft false
 * @link https://leetcode.cn/problems/first-unique-character-in-a-string/
 * @frontendId 387
 */

 func firstUniqChar(s string) int {
    cnt := [26]int{}
    for _, ch := range s {
        cnt[ch-'a']++
    }
    for i, ch := range s {
        if cnt[ch-'a'] == 1 {
            return i
        }
    }
    return -1
}