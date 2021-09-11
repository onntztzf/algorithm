package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

//167.两数之和 II - 输入有序数组
//link: https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		}

	}
	return []int{0, 0}
}
