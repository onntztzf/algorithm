/*
	66. 加一

	给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

	最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

	你可以假设除了整数 0 之外，这个整数不会以零开头。

	示例 1:
		输入: [1,2,3]
		输出: [1,2,4]
		解释: 输入数组表示数字 123。

	示例 2:
		输入: [4,3,2,1]
		输出: [4,3,2,2]
		解释: 输入数组表示数字 4321。
*/

package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}

func plusOne(digits []int) []int {
	//从后向前遍历
	for i := len(digits) - 1; i >= 0; i-- {
		//如果当前位不是 9，则对当前位加 1，然后返回原数组
		//如果当前位是 9，则将当前位变为 0
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	//如果数组中的元素都是 9，经过上面的操作，就会全部变为 0。然后我们在最前面补充一个 1
	return append([]int{1}, digits...)
}

func plusOne1(digits []int) []int {
	var result = make([]int, 0)
	var temp = make([]int, 0)
	carry := false
	for i := len(digits) - 1; i >= 0; i-- {
		current := digits[i]
		if i == len(digits) - 1 {
			current = current + 1
		}
		if carry {
			current += 1
		}
		if current >= 10 {
			carry = true
			current = current % 10
		} else {
			carry = false
		}
		temp = append(temp, current)
	}
	if carry {
		temp = append(temp, 1)
	}
	for i := len(temp) - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}
	return result
}
