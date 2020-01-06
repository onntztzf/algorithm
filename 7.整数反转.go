/*
	7.整数反转

	给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
	示例 1:
		输入: 123
		输出: 321
	示例 2:
		输入: -123
		输出: -321
	示例 3:
		输入: 120
		输出: 21
	注意:
	假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-2147483648))
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		//取模，获取最后一位数字
		last := x % 10
		//移除 x 的最后一位数
		x = x/10

		//题目中描述：环境只能存储得下 32 位的有符号整数
		//所以需要提前做校验

		//当 result * 10 > math.MaxInt32 一定越界
		//当 result * 10 = math.MaxInt32 时，如果对 x 取模后大于 7，则一定越界
		if result * 10 > math.MaxInt32 || ((result * 10 == math.MaxInt32) && last > 7) {
			return 0
		}

		//当 result * 10 < math.MinInt32 一定越界
		//当 result * 10 = math.MinInt32 时，如果对 x 取模后小于 -8，则一定越界
		if result * 10 < math.MinInt32 || ((result * 10 == math.MinInt32) && last < -8) {
			return 0
		}

		//result 添加一位，然后加上上面对 x 取模的结果
		result = result * 10 + last
	}
	return result
}