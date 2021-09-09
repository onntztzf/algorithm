package main

func main() {
}

//283.移动零
//link: https://leetcode-cn.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	for left, right, n := 0, 0, len(nums); right < n; {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
