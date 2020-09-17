/*
	58. 最后一个单词的长度

	给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

	如果不存在最后一个单词，请返回 0。

	说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。

	示例:
		输入: "Hello World"
		输出: 5
*/

package main

func main() {
	print(lengthOfLastWord("a "))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	//最后一个单词的长度
	var wordLength = 0
	//是否开始计数
	var startCounting = false
	for count := len(s) - 1; count >= 0; count-- {
		//当前字符为空格时
		if s[count] == ' ' {
			//如果已经开始计数，那么遇到空格则表示一个单词的结束，返回单词的长度
			//如果没有开始计数，说明从结尾开始是连续的空格，需要跳过，然后继续执行
			if startCounting == true {
				break
			} else {
				continue
			}
		}
		//当遇上不是空格的字符，开始进行计算单词长度
		startCounting = true
		wordLength++
	}
	return wordLength
}
