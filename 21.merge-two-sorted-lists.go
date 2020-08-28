/*
	21. 合并两个有序链表
	将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

	示例：
		输入：1->2->4, 1->3->4
		输出：1->1->2->3->4->4
*/
package main

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	mergeTwoLists(l1, l2)

	mergeTwoLists1(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//迭代
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}
	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}
	return result
}

//递归
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	if l1.Val < l2.Val{
//		l1.Next = mergeTwoLists(l1.Next,l2)
//		return l1
//	}else{
//		l2.Next = mergeTwoLists(l1,l2.Next)
//		return l2
//	}
//}
