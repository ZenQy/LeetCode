// Created by Bob at 2024/05/10 21:04
// leetgo: dev
// https://leetcode.cn/problems/minimum-moves-to-equal-array-elements/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minMoves(nums []int) (ans int) {
	min := nums[0]
	count := len(nums)
	for i := 0; i < count; i++ {
		if min > nums[i] {
			min = nums[i]
		}
		ans += nums[i]
	}
	ans -= min * count
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := minMoves(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
