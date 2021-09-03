/**
 * @brief
 * @file 88.merge-sorted-array
 * @author zhangpeng<zhangpeng.0304@aliyun.com>
 * @version 1.0
 * @date
 */

package main

//88. 合并两个有序数组
//https://leetcode-cn.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int)  {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		cur := 0
		if p1 < 0 {
			cur = nums2[p2]
			p2--
		} else if p2 < 0 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}