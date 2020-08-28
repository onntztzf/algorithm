/*
	53. 最大子序和

	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	示例:

		输入: [-2,1,-3,4,-1,2,1,-5,4]
		输出: 6
		解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

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
