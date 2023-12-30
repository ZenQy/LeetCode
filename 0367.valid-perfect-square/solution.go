// Created by ZenQy at 2023/12/30 22:20
// leetgo: dev
// https://leetcode.cn/problems/valid-perfect-square/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPerfectSquare(num int) bool {
	var mid, mul int
	min, max := 1, num
	for min <= max {
		mid = (min + max) / 2
		mul = mid * mid
		if mul > num {
			max = mid - 1
			continue
		}
		if mul < num {
			min = mid + 1
			continue
		}
		if mul == num {
			return true
		}
	}
	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num := Deserialize[int](ReadLine(stdin))
	ans := isPerfectSquare(num)

	fmt.Println("\noutput:", Serialize(ans))
}
