// Created by ZenQy at 2024/01/01 22:39
// leetgo: dev
// https://leetcode.cn/problems/squares-of-a-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sortedSquares(nums []int) []int {
	lenth := len(nums)
	ans := make([]int, lenth)
	for left, right, i := 0, lenth-1, lenth-1; i >= 0; i-- {
		if nums[left]+nums[right] > 0 {
			ans[i] = nums[right] * nums[right]
			right--
		} else {
			ans[i] = nums[left] * nums[left]
			left++
		}
	}
	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := sortedSquares(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
