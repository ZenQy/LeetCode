package solve0451

import "bytes"

/**
 * @index 451
 * @title 根据字符出现频率排序
 * @difficulty 中等
 * @tags hash-table,string,bucket-sort,counting,sorting,heap-priority-queue
 * @draft false
 * @link https://leetcode.cn/problems/sort-characters-by-frequency/
 * @frontendId 451
 */

func frequencySort(s string) string {
	cnt := map[byte]int{}
	maxFreq := 0
	for i := range s {
		cnt[s[i]]++
		maxFreq = max(maxFreq, cnt[s[i]])
	}

	buckets := make([][]byte, maxFreq+1)
	for ch, c := range cnt {
		buckets[c] = append(buckets[c], ch)
	}

	ans := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
