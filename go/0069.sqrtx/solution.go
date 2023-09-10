// Created by ZenQy at 2023/09/10 18:08
// leetgo: dev
// https://leetcode.cn/problems/sqrtx/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mySqrt(x int) int {
	left, mid, right := 0, 0, x
	for left <= right {
		mid = (left + right) >> 1
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := mySqrt(x)

	fmt.Println("\noutput:", Serialize(ans))
}
