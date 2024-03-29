// Created by ZenQy at 2023/12/30 22:17
// leetgo: dev
// https://leetcode.cn/problems/remove-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeElement(nums []int, val int) (ans int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ans] = nums[i]
			ans++
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	val := Deserialize[int](ReadLine(stdin))
	ans := removeElement(nums, val)

	fmt.Println("\noutput:", Serialize(ans))
}
