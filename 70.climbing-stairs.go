package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

//70. 爬楼梯
//link: https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
