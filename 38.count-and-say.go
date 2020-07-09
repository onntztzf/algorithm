/*
	38. 外观数列

	给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。

	注意：整数序列中的每一项将表示为一个字符串。

	「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
		1.     1
		2.     11
		3.     21
		4.     1211
		5.     111221

	第一项是数字 1

	描述前一项，这个数是 1 即 “一个 1 ”，记作 11

	描述前一项，这个数是 11 即 “两个 1 ” ，记作 21

	描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211

	描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221

	示例 1:
		输入: 1
		输出: "1"
		解释：这是一个基本样例。

	示例 2:
		输入: 4
		输出: "1211"
		解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(3))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	//上一次的计算结果
	var lastResult = countAndSay(n-1)
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
