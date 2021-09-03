package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(3))
}

//38. 外观数列
//link: https://leetcode-cn.com/problems/implement-strstr/
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	//上一次的计算结果
	var lastResult = countAndSay(n - 1)
	//本次的结果
	var currentResult = ""
	//相同字符的计数器
	var count = 1
	//对上一次(即n-1)的结果进行遍历
	for i := 0; i <= len(lastResult)-1; i++ {
		//如果当前位是最后一位，直接拼接
		if i == len(lastResult)-1 {
			currentResult += fmt.Sprintf("%+v%s", count, string(lastResult[i]))
			return currentResult
		}
		//如果当前位不是最后一位
		if lastResult[i] != lastResult[i+1] {
			//当前位和下一位不相同，拼接后，字符计数器重置
			currentResult += fmt.Sprintf("%+v%s", count, string(lastResult[i]))
			count = 1
		} else {
			//当前位和下一位相同，字符计数器+1
			count++
		}
	}
	return currentResult
}
