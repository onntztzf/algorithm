/**
 * @brief
 * @file 977.squares-of-a-sorted-array
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
}
//977. 有序数组的平方
//link: https://leetcode-cn.com/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[left] * nums[left] > nums[right] * nums[right] {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
	}
	return result
}