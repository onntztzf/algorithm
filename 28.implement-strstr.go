/*
	28. 实现 strStr()

	实现 strStr() 函数。

	给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

	示例 1:

		输入: haystack = "hello", needle = "ll"
		输出: 2

	示例 2:

		输入: haystack = "aaaaa", needle = "bba"
		输出: -1

	说明:

		当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

		对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/
package main

import "fmt"

func main() {
	//fmt.Println(strStr("a", "a"))
	//fmt.Println(strStr("hello", "ll"))
	//fmt.Println(strStr("mississippi", "issipi"))
	//fmt.Println(strStr("mississippi", "mississippi"))
	fmt.Println(strStr("mississippi", "pi"))
}

func strStr(haystack string, needle string) int {
	ln := len(needle)
	if ln == 0 {
		return 0
	}
	lh := len(haystack)
	if ln > lh {
		return -1
	}
	for i := 0; i <= lh - ln; i++ {
		sameLength := 0
		for j := 0; j < ln; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			sameLength++
		}
		if sameLength == ln {
			return i
		}
	}
	return -1
}