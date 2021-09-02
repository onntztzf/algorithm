/**
 * @brief
 * @file 278.first-bad-version
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package main

func main()  {
	
}

//278. 第一个错误的版本
//link: https://leetcode-cn.com/problems/first-bad-version/
func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right - left) / 2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}