// Created by Bob at 2024/06/27 21:59
// leetgo: dev
// https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findDisappearedNumbers(nums []int) (ans []int) {
	for i := 0; i < len(nums); i++ {
		n:=abs(nums[i])-1
		nums[n] = -abs(nums[n])
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findDisappearedNumbers(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
