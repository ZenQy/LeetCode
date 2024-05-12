// Created by Bob at 2024/05/12 20:56
// leetgo: dev
// https://leetcode.cn/problems/rotate-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	count := len(nums)
	for i := 0; i < count/2; i++ {
		nums[i], nums[count-1-i] = nums[count-1-i], nums[i]
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	rotate(nums, k)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
