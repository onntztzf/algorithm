/*
	67. 二进制求和

	给你两个二进制字符串，返回它们的和（用二进制表示）。

	输入为 非空 字符串且只包含数字 1 和 0。

	示例 1:
		输入: a = "11", b = "1"
		输出: "100"

	示例 2:
		输入: a = "1010", b = "1011"
		输出: "10101"

	提示：
		每个字符串仅由字符 '0' 或 '1' 组成。
		1 <= a.length, b.length <= 10^4
		字符串如果不是 "0" ，就都不含前导零。
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("10001", "1"))
}

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	len := lenA
	if lenB > lenA{
		len = lenB
	}

	for i := 0; i < len; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		result = strconv.Itoa(carry%2) + result
		carry /= 2
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}
