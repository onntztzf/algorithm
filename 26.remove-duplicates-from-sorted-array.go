package main

func main() {
}

//26. 删除排序数组中的重复项
//link: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	left, right := 1, 1
	for right < n {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
