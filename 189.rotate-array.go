package main

func main() {
	rotate([]int{1, 2, 3, 4, 5}, 3)
}

//旋转数组
//link: https://leetcode-cn.com/problems/rotate-array
func rotate(nums []int, k int) {
	//nums = "----->-->"; k =3
	//result = "-->----->";

	//reverse "----->-->" we can get "<--<-----"
	//reverse "<--" we can get "--><-----"
	//reverse "<-----" we can get "-->----->"
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
