package main

func main() {
	//longestCommonPrefix1([]string{"flower","flow","flight"})
}

//14. 最长公共前缀
//link: https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//取第一项当作公共前缀
	var commonPrefix = strs[0]
	//找到最短的字符串当作公共前缀
	for _, val := range strs {
		if len(commonPrefix) > len(val) {
			commonPrefix = val
		}
	}
	//遍历字符串数组
	for _, val := range strs {
		//获取当前字符串和 commonPrefix 长度一样的前缀 currentPrefix
		currentPrefix := val[0:len(commonPrefix)]
		//如果 currentPrefix 和 commonPrefix 不一样，则进入下面循环
		for currentPrefix != commonPrefix {
			//删除 commonPrefix 最后一位
			commonPrefix = commonPrefix[0 : len(commonPrefix)-1]
			//当 commonPrefix 修改为空字符串时，表示没有公共前缀，返回 ""
			if len(commonPrefix) == 0 {
				return ""
			}
		}
	}
	//返回公共字符串
	return commonPrefix
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//取第一项当作公共前缀
	var commonPrefix = strs[0]
	//遍历公共前缀
	for i := 0; i < len(commonPrefix); i++ {
		//遍历字符串数组
		for j := 1; j < len(strs); j++ {
			//公共前缀的当前下标 i 和当前的字符串长度相同
			if i == len(strs[j]) {
				return commonPrefix[0:i]
			}
			//公共前缀的当前字符
			currentCommonPrefixCharacter := commonPrefix[i]
			//当前字符串的当前字符
			currentCharacter := strs[j][i]
			//如果 currentCommonPrefixCharacter 与 currentCharacter 不同
			if currentCommonPrefixCharacter != currentCharacter {
				return commonPrefix[0:i]
			}
		}
	}
	return commonPrefix
}
