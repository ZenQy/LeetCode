// Created by Bob at 2024/05/12 19:48
// leetgo: dev
// https://leetcode.cn/problems/non-decreasing-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func checkPossibility(nums []int) bool {
	n := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			n++
			if i > 0 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			}
		}
		if n > 1 {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := checkPossibility(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
