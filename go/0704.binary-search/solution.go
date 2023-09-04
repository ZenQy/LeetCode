// Created by ZenQy at 2023/08/30 11:59
// leetgo: dev
// https://leetcode.cn/problems/binary-search/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func search(nums []int, target int) (ans int) {
	left, right := 0, len(nums)
	for left < right {
		ans = (left + right) / 2
		if nums[ans] == target {
			return
		}
		if nums[ans] < target {
			left = ans+1
		} else {
			right = ans
		}
	}
	ans = -1
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := search(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
