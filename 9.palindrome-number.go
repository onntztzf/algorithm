/*
	9.回文数

	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

	示例 1:
		输入: 121
		输出: true
	示例 2:
		输入: -121
		输出: false
		解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
	示例 3:
		输入: 10
		输出: false
		解释: 从右向左读, 为 01 。因此它不是一个回文数。

	进阶:
		你能不将整数转为字符串来解决这个问题吗？
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome2(1234321))
}

//9. 回文数
//link: https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	//小于 0 || 尾数为 0
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	//一位数都是回文数
	if x < 10 {
		return true
	}
	//反转结果
	reverseResult := 0
	for x > reverseResult {
		reverseResult = reverseResult*10 + x%10
		x /= 10
	}
	//第一个判断为整数长度为偶数时，第二个判断为整数长度为奇数时（奇数时 中间那一位肯定会与自己相等）
	return x == reverseResult || x == reverseResult/10
}

func isPalindrome2(x int) bool {
	string := strconv.Itoa(x)
	for i := 0; i < len(string)/2; i++ {
		if string[i] != string[len(string)-1-i] {
			return false
		}
	}
	return true
}
