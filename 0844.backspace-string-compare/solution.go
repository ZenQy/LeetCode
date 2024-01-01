// Created by ZenQy at 2024/01/01 21:54
// leetgo: dev
// https://leetcode.cn/problems/backspace-string-compare/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func backspaceCompare(s string, t string) bool {
	return realString(s) == realString(t)
}

func realString(s string) string {
	bs := []byte(s)
	hash := []byte("#")[0]
	low := 0

	for i := 0; i < len(bs); i++ {
		if bs[i] == hash {
			if low > 0 {
				low--
			}
		} else {
			bs[low] = bs[i]
			low++
		}
	}

	return string(bs[:low])
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := backspaceCompare(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
