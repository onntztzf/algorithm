package main

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
}

//27. 移除元素
//link: https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	n := len(nums)
	left, right := 0, 0
	for right < n {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
