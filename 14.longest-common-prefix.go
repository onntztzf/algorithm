/*
	14.最长公共前缀

	编写一个函数来查找字符串数组中的最长公共前缀。

    如果不存在公共前缀，返回空字符串 ""。

    示例 1:
		输入: ["flower","flow","flight"]
        输出: "fl"
    示例 2:
		输入: ["dog","racecar","car"]
        输出: ""
        解释: 输入不存在公共前缀。
   说明:
	所有输入只包含小写字母 a-z。
*/
package main

func main() {
	longestCommonPrefix([]string{"flower","flow","flight"})
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var commonPrefix = strs[0]
	for _, val := range strs {
		if len(commonPrefix) > len(val) {
			commonPrefix = val
		}
	}
	for _, val := range strs {
		for val[0:len(commonPrefix)] != commonPrefix {
			commonPrefix = commonPrefix[0:len(commonPrefix)-1]
			if len(commonPrefix) == 0 {
				return ""
			}
		}
	}
	return commonPrefix
}
