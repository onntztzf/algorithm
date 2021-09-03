package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IXI"))
}

//13. 罗马数字转整数
//link: https://leetcode-cn.com/problems/roman-to-integer/
func romanToInt(s string) int {
	m := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if _, ok := m[s[i:i+2]]; ok {
				result += m[s[i:i+2]]
				i++
				continue
			}
		}
		if _, ok := m[s[i:i+1]]; ok {
			result += m[s[i:i+1]]
		}
	}
	return result
}
