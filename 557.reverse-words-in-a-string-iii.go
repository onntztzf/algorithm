package main

import "fmt"

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

//557. 反转字符串中的单词 III
//link: https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
func reverseWords(s string) string {
	length := len(s)
	var ret []byte
	for i := 0; i < length; i++ {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			ret = append(ret, s[start + i - 1 - p])
		}
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}