package main

func main() {
}

//344. 反转字符串
//link: https://leetcode-cn.com/problems/reverse-string
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}
