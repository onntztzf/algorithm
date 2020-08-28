package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return idx
		} else {
			idx++
		}
	}
	return idx
}
