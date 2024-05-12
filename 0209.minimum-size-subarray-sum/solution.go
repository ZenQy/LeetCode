// Created by ZenQy at 2024/01/03 21:32
// leetgo: dev
// https://leetcode.cn/problems/minimum-size-subarray-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minSubArrayLen(target int, nums []int) (ans int) {
	sum := 0
	lenth := len(nums)
	ans = lenth + 1
	for i, j := 0, 0; j < lenth; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < ans {
				ans = subLength
			}
			sum -= nums[i]
			i++
		}
	}

	if ans > lenth {
		ans = 0
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	target := Deserialize[int](ReadLine(stdin))
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := minSubArrayLen(target, nums)

	fmt.Println("\noutput:", Serialize(ans))
}
