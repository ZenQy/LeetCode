package solve0274

import (
	"fmt"
	"sort"
)

/**
 * @index 274
 * @title H 指数
 * @difficulty 中等
 * @tags array,counting-sort,sorting
 * @draft false
 * @link https://leetcode.cn/problems/h-index/
 * @frontendId 274
 */

func hIndex(citations []int) int {
	sort.Ints(citations)
	fmt.Println(citations)
	ln := len(citations)
	for i := 1; i <= ln; i++ {
		if citations[ln-i] < i {
			return i-1
		}
	}
	return ln
}
