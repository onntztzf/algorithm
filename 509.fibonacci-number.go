package main

import "fmt"

/**
509. 斐波那契数

斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1

给定 n ，请计算 F(n) 。



示例 1：

输入：n = 2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1

示例 2：

输入：n = 3
输出：2
解释：F(3) = F(2) + F(1) = 1 + 1 = 2

示例 3：

输入：n = 4
输出：3
解释：F(4) = F(3) + F(2) = 2 + 1 = 3

提示：

    0 <= n <= 30
*/

func main() {
	fmt.Printf("dp table fib result: %d\n", fib2(5))
	fmt.Printf("map cache fib result: %d\n", fib1(5))
}

//通过dp table提高效率
func fib2(n int) int {
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

//通过缓存结果提高效率
func fib1(n int) int {
	temp := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}
	f := func(n int) int {
		if v, ok := temp[n]; ok {
			return v
		}
		temp[n] = fib1(n-1) + fib1(n-2)
		return temp[n]
	}
	return f(n)
}
