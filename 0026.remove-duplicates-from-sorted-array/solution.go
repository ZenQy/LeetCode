// Created by ZenQy at 2023/12/31 22:50
// leetgo: dev
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeDuplicates(nums []int) (ans int) {
	for i := 1; i < len(nums); i++ {
		if nums[ans] != nums[i] {
			ans++
			nums[ans] = nums[i]
		}
	}
	ans++
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := removeDuplicates(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
