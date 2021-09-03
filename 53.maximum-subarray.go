package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

//53. 最大子序和
//link: https://leetcode-cn.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	//取数组第 0 项当作最大值
	max := nums[0]
	//从数组第 1 项开始遍历
	for i := 1; i < len(nums); i++ {
		//如果上一项对当前项有增益效果，则将当前项加上上一项
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		//如果当前项大于最大值，则将最大值换为当前项
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
