/**
 * @brief
 * @file 704.binary-search
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package main

func main() {

}

func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target{
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}
