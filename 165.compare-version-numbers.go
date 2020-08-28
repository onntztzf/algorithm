/*
	165.比较版本号

	比较两个版本号 version1 和 version2。
	如果 version1 > version2 返回 1，如果 version1 < version2 返回 -1， 除此之外返回 0。

	你可以假设版本字符串非空，并且只包含数字和 . 字符。

	. 字符不代表小数点，而是用于分隔数字序列。

	例如，2.5 不是“两个半”，也不是“差一半到三”，而是第二版中的第五个小版本。

	你可以假设版本号的每一级的默认修订版号为 0。例如，版本号 3.4 的第一级（大版本）和第二级（小版本）修订号分别为 3 和 4。其第三级和第四级修订号均为 0。
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.5.2.6", "1.5.2.6"))
}

func compareVersion(version1 string, version2 string) int {
	//以 . 对版本号进行分割
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")

	//获取最大的版本号长度
	maxLength := math.Max(float64(len(arr1)), float64(len(arr2)))

	for index := 0; index < int(maxLength); index++ {
		v1 := ""
		v2 := ""

		//获取第 index 位的版本号
		if index < len(arr1) {
			v1 = arr1[index]
		}
		if index < len(arr2) {
			v2 = arr2[index]
		}

		//将字符串转为 int
		i1, _ := strconv.Atoi(v1)
		i2, _ := strconv.Atoi(v2)

		if i1 < i2 {
			return -1
		} else if i1 > i2 {
			return 1
		}
	}
	return 0
}
