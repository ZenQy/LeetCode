// Created by ZenQy at 2023/12/31 23:04
// leetgo: dev
// https://leetcode.cn/problems/move-zeroes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func moveZeroes(nums []int) {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[ans] = nums[i]
			ans++
		}
	}
	for i := ans; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	moveZeroes(nums)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
