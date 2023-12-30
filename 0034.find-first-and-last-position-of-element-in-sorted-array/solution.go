// Created by ZenQy at 2023/08/31 16:50
// leetgo: dev
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchRange(nums []int, target int) (ans []int) {
	ans = []int{-1, -1}
	if len(nums) == 0 {
		return
	}
	left := Boader(nums, target, false)
	if len(nums) == left {
		return
	}
	if nums[left] != target {
		return
	}
	ans[0] = left
	ans[1] = Boader(nums, target, true) - 1

	return
}

func Boader(nums []int, target int, more bool) int {
	left, mid, right := 0, 0, len(nums)
	for left < right {
		mid = (left + right) / 2
		if nums[mid] < target || (nums[mid] == target && more) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchRange(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
