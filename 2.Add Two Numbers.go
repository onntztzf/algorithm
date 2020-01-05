/*
	1.两数相加

	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

    如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

    您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	示例:
		输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
        输出：7 -> 0 -> 8
        原因：342 + 465 = 807
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	output := addTwoNumbers(l1, l2)

	for true {
		fmt.Println(output.Val)
		output = output.Next
		if output == nil {
			break
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//return enumerate(l1, l2)
	//return enumerate2(l1, l2)
	return recurse(l1, l2)
}

/*
遍历
*/
func enumerate(l1 *ListNode, l2 *ListNode) *ListNode {
	//创建和的链表
	resultList := &ListNode{}

	//获取第一个节点
	currentNode := resultList

	//进位标志
	carry := 0

	//当 l1 或 l2 不为空时，进行遍历操作
	for l1 != nil || l2 != nil {
		//创建 x, y 保存 l1 和 l2 当前节点的 Val
		x, y := 0, 0
		//l1 当前节点不为空时，对 x 赋值, 同时修改 l1 指向的节点
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		//l2 当前节点不为空时，对 x 赋值, 同时修改 l2 指向的节点
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		//对 x, y, carry 求和
		sum := x + y + carry
		//计算是否需要进位
		carry = sum / 10
		//resultList 的当前节点 Val 进行赋值
		currentNode.Val = sum % 10

		//当 l1 或 l2 当前节点不为空，或者应该进位时，为 resultList 创建下一个节点
		if l1 != nil || l2 != nil || carry != 0 {
			currentNode.Next = &ListNode{carry, nil}
		}

		//本次计算完毕后，将 currentNode 修改为下一个节点
		currentNode = currentNode.Next
	}
	return resultList
}

func enumerate2(l1 *ListNode, l2 *ListNode) *ListNode {
	//创建和的链表
	resultList := &ListNode{}
	//获取第一个节点
	currentNode := resultList
	//进位标志
	carry := 0

	//当 l1 或 l2 不为空时，进行遍历操作
	for l1 != nil || l2 != nil {
		//创建 x, y 保存 l1 和 l2 当前节点的 Val
		x, y := 0, 0
		//l1 当前节点不为空时，对 x 赋值, 同时修改 l1 指向的节点
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		//l2 当前节点不为空时，对 x 赋值, 同时修改 l2 指向的节点
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		//对 x, y, carry 求和
		sum := x + y + carry
		//计算是否需要进位
		carry = sum / 10
		currentNode.Next = &ListNode{sum % 10, nil}
		//本次计算完毕后，将 currentNode 修改为下一个节点
		currentNode = currentNode.Next
	}
	if carry != 0 {
		currentNode.Next = &ListNode{carry, nil}
	}
	return resultList.Next
}

/*
递归
*/
func recurse(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	nextNode := addTwoNumbers(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{ Val: sum, Next: nextNode }
	} else {
		tempNode := &ListNode{
			Val: 1,
			Next: nil,
		}
		return &ListNode{
			Val: sum - 10,
			Next: addTwoNumbers(nextNode, tempNode),
		}
	}
}
