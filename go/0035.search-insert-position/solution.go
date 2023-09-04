// Created by ZenQy at 2023/08/31 16:34
// leetgo: dev
// https://leetcode.cn/problems/search-insert-position/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchInsert(nums []int, target int) (ans int) {
	left, right := 0, len(nums)
	for left < right {
		ans = (left + right) / 2
		if nums[ans] == target {
			return ans
		}
		if nums[ans] > target {
			right = ans
		} else {
			left = ans + 1
		}
	}
	if nums[ans] < target {
		ans++
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchInsert(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
