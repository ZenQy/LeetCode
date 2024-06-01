// Created by Bob at 2024/06/01 10:52
// leetgo: dev
// https://leetcode.cn/problems/set-mismatch/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findErrorNums(nums []int) []int {
	count := len(nums)
	ans := make([]int, 2)
	x := 0
	for i := 0; i < count; i++ {
		x = abs(nums[i])
		if nums[x-1] < 0 {
			ans[0] = x
		} else {
			nums[x-1] *= -1
		}
	}
	for i := 0; i < count; i++ {
		if nums[i] > 0 {
			ans[1] = i + 1
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findErrorNums(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
