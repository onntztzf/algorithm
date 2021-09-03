package main

import "fmt"

func main() {
	//fmt.Println(strStr("a", "a"))
	//fmt.Println(strStr("hello", "ll"))
	//fmt.Println(strStr("mississippi", "issipi"))
	//fmt.Println(strStr("mississippi", "mississippi"))
	fmt.Println(strStr("mississippi", "pi"))
}

//28. 实现 strStr()
//link: https://leetcode-cn.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	ln := len(needle)
	if ln == 0 {
		return 0
	}
	lh := len(haystack)
	if ln > lh {
		return -1
	}
	for i := 0; i <= lh-ln; i++ {
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
