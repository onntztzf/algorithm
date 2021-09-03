package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("10001", "1"))
}

//67. 二进制求和
//link: https://leetcode-cn.com/problems/add-binary/
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
