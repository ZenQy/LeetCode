// Created by Bob at 2024/05/13 20:07
// leetgo: dev
// https://leetcode.cn/problems/rotate-function/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxRotateFunction(nums []int) (ans int) {
	numSum := 0
	count := len(nums)
	f := 0
	for i := 0; i < count; i++ {
		numSum += nums[i]
		f += nums[i] * i
	}

	ans = f
	for i := 1; i < count; i++ {
		f += numSum - nums[count-i]*count
		if f > ans {
			ans = f
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxRotateFunction(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
