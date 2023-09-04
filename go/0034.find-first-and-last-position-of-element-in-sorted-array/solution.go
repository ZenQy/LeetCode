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

	left, mid, right := 0, 0, len(nums)
	if right == 0 {
		return
	}
	for left < right {
		mid = (left + right) / 2
		fmt.Println(mid, nums[mid])
		if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
			break
		}
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[mid] != target {
		return
	}
	ans[0] = mid

	left, mid, right = 0, 0, len(nums)
	for left < right {
		mid = (left + right) / 2
		fmt.Println(mid, nums[mid])
		if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] > target) {
			break
		}
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	ans[1] = mid
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchRange(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
